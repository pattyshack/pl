package build

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/analyze"
	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/build/cfg"
	"github.com/pattyshack/pl/errors"
	"github.com/pattyshack/pl/parser"
)

type BuildMode int

const (
	BuildLibrary = BuildMode(0)
	BuildBinary  = BuildMode(1) // lib + bin
	BuildTest    = BuildMode(2) // lib + bin + test
)

type pkgLoc struct {
	*Package
	lexutil.Location
}

type locateState struct {
	*errors.Emitter

	mutex sync.Mutex

	// All guarded by mutex
	located       bool
	err           error
	isBuildTarget bool
	importedBy    []pkgLoc
}

func (state *locateState) Locate(ws *Workspace, id ast.PackageID) (
	PackageContents,
	error,
) {
	state.mutex.Lock()
	defer state.mutex.Unlock()

	if state.located {
		panic("should never happen")
	}
	state.located = true

	contents, err := ws.Locate(id)
	state.err = err

	if err != nil {
		if state.isBuildTarget {
			state.EmitErrors(err)
		}

		for _, by := range state.importedBy {
			by.Package.EmitErrors(lexutil.LocationError{Loc: by.Location, Err: err})
		}
	}

	return contents, err
}

func (state *locateState) ImportedBy(by *pkgLoc) {
	state.mutex.Lock()
	defer state.mutex.Unlock()

	if !state.located {
		if by == nil {
			state.isBuildTarget = true
		} else {
			state.importedBy = append(state.importedBy, *by)
		}
		return
	}

	if by == nil {
		if state.isBuildTarget {
			return // already emitted error, if any
		}
		state.isBuildTarget = true

		if state.err != nil {
			state.EmitErrors(state.err)
		}
	} else if state.err != nil {
		by.Package.EmitErrors(
			lexutil.LocationError{
				Loc: by.Location,
				Err: state.err,
			})
	}
}

type Package struct {
	builder *Builder

	BuildMode

	ast.PackageID

	*errors.Emitter
	locateState

	// Populated by Load().  Safe to access once loadWaitGroup is ready
	PackageContents

	LibraryDefinitions *ast.StatementList
	TestDefinitions    *ast.StatementList

	DirectDependencies map[ast.PackageID]pkgLoc

	// Only used by detectCycles, after loadWaitGroup is ready, and
	// sortedPackages is available.
	isAcyclic bool

	PublicInterfaceAnalyzed chan struct{}
}

func (pkg *Package) Load() error {
	defer pkg.builder.loadWaitGroup.Done()

	// TODO load cached entry

	contents, err := pkg.locateState.Locate(pkg.builder.Workspace, pkg.PackageID)
	if err != nil {
		return err
	}
	pkg.PackageContents = contents

	parseOptions := parser.ParserOptions{
		OpenFunc: contents.Open,
	}

	libSrcs, testSrcs := contents.Sources()
	if pkg.BuildMode == BuildLibrary {
		testSrcs = nil
	} else if pkg.BuildMode == BuildBinary {
		testSrcs = nil
	}

	allSrcs := make([]string, 0, len(libSrcs)+len(testSrcs))
	allSrcs = append(allSrcs, libSrcs...)
	allSrcs = append(allSrcs, testSrcs...)

	parsedSrcs := parser.ParseSources(
		allSrcs,
		pkg.builder.Config,
		pkg.Emitter,
		parseOptions)

	libDefs := ast.NewStatementList()
	testDefs := ast.NewStatementList()

	for idx, parsed := range parsedSrcs {
		if parsed.FailedBuildConstraints {
			continue
		}

		if idx < len(libSrcs) {
			libDefs.Elements = append(
				libDefs.Elements,
				parsed.Definitions.Elements...)
		} else {
			testDefs.Elements = append(
				testDefs.Elements,
				parsed.Definitions.Elements...)
		}

		for _, importPkg := range parsed.Dependencies {
			loc := importPkg.Loc()
			pkg.DirectDependencies[importPkg.PackageID] = pkgLoc{
				Package: pkg.builder.build(
					BuildLibrary,
					importPkg.PackageID,
					&pkgLoc{pkg, loc}),
				Location: loc,
			}
		}
	}

	pkg.LibraryDefinitions = libDefs
	pkg.TestDefinitions = testDefs
	return nil
}

func (pkg *Package) Build() {
	defer pkg.builder.buildWaitGroup.Done()
	defer close(pkg.PublicInterfaceAnalyzed)

	loadErr := pkg.Load()
	if loadErr != nil || pkg.HasErrors() {
		// TODO populate stub empty public interface
		return
	}

	for _, dep := range pkg.DirectDependencies {
		select {
		case <-pkg.builder.ctx.Done():
			return
		case <-dep.PublicInterfaceAnalyzed:
			// ok
		}
	}

	analyze.Semantic(pkg.LibraryDefinitions, pkg.Emitter)
}

func (pkg *Package) detectCycle(visited []ast.PackageID) error {
	for idx, visitedID := range visited {
		if visitedID == pkg.PackageID {
			cycle := visited[idx:]

			cycleErrMsg := "detected import dependency cycle:\n  "
			for idx, pkgId := range cycle {
				curr := pkg.builder.get(pkgId)
				cycleErrMsg += fmt.Sprintf(
					"%s (%s) ->\n  ",
					pkgId,
					curr.DirectDependencies[cycle[(idx+1)%len(cycle)]].Location)
			}
			cycleErrMsg += cycle[0].String()

			return fmt.Errorf(cycleErrMsg)
		}
	}

	visited = append(visited, pkg.PackageID)
	// Walk deps in deterministic order
	for _, dep := range pkg.builder.Packages() {
		if dep.isAcyclic {
			continue
		}

		_, ok := pkg.DirectDependencies[dep.PackageID]
		if !ok {
			continue
		}

		err := dep.detectCycle(visited)
		if err != nil {
			return err
		}
	}
	visited = visited[:len(visited)-1]

	pkg.isAcyclic = true
	return nil
}

type Builder struct {
	ctx    context.Context
	cancel context.CancelFunc

	*cfg.Config
	*Workspace

	*errors.Emitter

	loadWaitGroup  sync.WaitGroup
	buildWaitGroup sync.WaitGroup

	mutex sync.Mutex

	packages map[ast.PackageID]*Package // guarded by mutex.

	// Computed only after loadWaitGroup is ready.
	sortedPackages []*Package // guarded by mutex
}

func NewBuilder(
	workspace *Workspace,
	config *cfg.Config,
) *Builder {
	ctx, cancel := context.WithCancel(context.Background())
	builder := &Builder{
		ctx:       ctx,
		cancel:    cancel,
		Config:    config,
		Workspace: workspace,
		Emitter:   &errors.Emitter{},
		packages:  map[ast.PackageID]*Package{},
	}

	return builder
}

func (builder *Builder) pkgs() []*Package {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()

	pkgs := make([]*Package, 0, len(builder.packages))
	for _, pkg := range builder.packages {
		pkgs = append(pkgs, pkg)
	}

	return pkgs
}

func (builder *Builder) get(id ast.PackageID) *Package {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()

	return builder.packages[id]
}

func (builder *Builder) build(
	mode BuildMode,
	id ast.PackageID,
	by *pkgLoc,
) *Package {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()

	pkg := builder.packages[id]
	if pkg == nil {
		emitter := &errors.Emitter{}
		pkg = &Package{
			builder:   builder,
			BuildMode: mode,
			PackageID: id,
			Emitter:   emitter,
			locateState: locateState{
				Emitter: emitter,
			},
			DirectDependencies:      map[ast.PackageID]pkgLoc{},
			PublicInterfaceAnalyzed: make(chan struct{}),
		}
		builder.packages[id] = pkg

		builder.loadWaitGroup.Add(1)
		builder.buildWaitGroup.Add(1)
		go pkg.Build()
	}

	pkg.locateState.ImportedBy(by)
	return pkg
}

func (builder *Builder) setSorted(sorted []*Package) {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()

	builder.sortedPackages = sorted
}

func (builder *Builder) Packages() []*Package {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()

	return builder.sortedPackages
}

func (builder *Builder) Build(mode BuildMode, ids ...ast.PackageID) []error {
	for _, id := range ids {
		builder.build(mode, id, nil)
	}

	builder.loadWaitGroup.Wait()

	builder.topoSortPkgsAndDetectCycles()

	builder.buildWaitGroup.Wait()

	errs := builder.Errors()
	for _, pkg := range builder.Packages() {
		errs = append(errs, pkg.Errors()...)
	}

	return errs
}

func (builder *Builder) topoSortPkgsAndDetectCycles() {
	toCheck := builder.pkgs()
	numPkgs := len(toCheck)

	// This ensures topo sort will also be name sorted for independent packages
	sort.Slice(toCheck, func(i int, j int) bool {
		return toCheck[i].PackageID.String() < toCheck[j].PackageID.String()
	})

	var cycleFound *Package
	processed := map[*Package]bool{}
	sorted := make([]*Package, 0, numPkgs)
	for len(toCheck) > 0 {
		remaining := make([]*Package, 0, len(toCheck))
		for _, pkg := range toCheck {
			processedCount := 0
			for _, dep := range pkg.DirectDependencies {
				if processed[dep.Package] {
					processedCount++
				}
			}

			if processedCount == len(pkg.DirectDependencies) {
				if cycleFound == nil {
					pkg.isAcyclic = true
				}

				processed[pkg] = true
				sorted = append(sorted, pkg)
			} else {
				remaining = append(remaining, pkg)
			}
		}

		if len(toCheck) == len(remaining) {
			// Processed all non-cycles.  Arbitrarily pick the last package to break
			// the cycle.
			lastIdx := len(toCheck) - 1
			pkg := toCheck[lastIdx]
			processed[pkg] = true
			sorted = append(sorted, pkg)

			if cycleFound == nil {
				cycleFound = pkg
			}

			toCheck = toCheck[:lastIdx]
		} else {
			toCheck = remaining
		}
	}

	builder.setSorted(sorted)

	if cycleFound != nil {
		builder.cancel()

		// NOTE: There could be multiple import cycles, but for simplicity, we'll
		// only report the first detected cycle.
		err := cycleFound.detectCycle(nil)
		if err == nil {
			panic("This should never happen")
		}
		builder.EmitErrors(err)
	}
}
