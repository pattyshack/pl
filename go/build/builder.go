package build

import (
	"context"
	"fmt"
	"sync"

	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/errors"
	"github.com/pattyshack/pl/parser"
	"github.com/pattyshack/pl/types"
)

type importedBy struct {
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
	importedBy    []importedBy
}

func (state *locateState) Locate(ws *Workspace, id types.PackageID) (
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

func (state *locateState) ImportedBy(by *importedBy) {
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

type depPkg struct {
	lexutil.Location
	*Package
}

type Package struct {
	builder *Builder

	types.PackageID

	*errors.Emitter
	locateState

	// Populated by Load().  Safe to access once loadWaitGroup is ready
	PackageContents
	Definitions        *ast.StatementList
	DirectDependencies map[types.PackageID]depPkg

	// Only used by detectCycles, after loadWaitGroup is ready.
	isAcyclic bool

	PublicInterfaceAnalyzed chan struct{}
}

func (pkg *Package) Load() {
	defer pkg.builder.loadWaitGroup.Done()

	// TODO load cached entry

	contents, err := pkg.locateState.Locate(pkg.builder.Workspace, pkg.PackageID)
	if err != nil {
		return
	}
	pkg.PackageContents = contents

	parseOptions := parser.ParserOptions{
		OpenFunc: contents.Open,
	}

	definitions, importPkgs := parser.ParsePackage(
		contents.Sources(),
		pkg.Emitter,
		parseOptions)
	pkg.Definitions = definitions

	for _, importPkg := range importPkgs {
		loc := importPkg.Loc()
		pkg.DirectDependencies[importPkg.PackageID] = depPkg{
			Location: loc,
			Package:  pkg.builder.build(importPkg.PackageID, &importedBy{pkg, loc}),
		}
	}
}

func (pkg *Package) Build() {
	defer pkg.builder.buildWaitGroup.Done()
	defer close(pkg.PublicInterfaceAnalyzed)

	pkg.Load()
	// TODO need to check for error / early exit

	for _, dep := range pkg.DirectDependencies {
		select {
		case <-pkg.builder.ctx.Done():
			return
		case <-dep.PublicInterfaceAnalyzed:
			// ok
		}
	}

	// can begin semantic analysis
}

func (pkg *Package) detectCycle(visited []types.PackageID) error {
	for idx, visitedID := range visited {
		if visitedID == pkg.PackageID {
			cycle := visited[idx:]

			cycleErrMsg := "detected import dependency cycle:\n  "
			for idx, pkgId := range cycle {
				curr := pkg.builder.packages[pkgId]
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
	for _, dep := range pkg.DirectDependencies {
		if dep.isAcyclic {
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

	*errors.Emitter

	loadWaitGroup  sync.WaitGroup
	buildWaitGroup sync.WaitGroup

	*Workspace

	mutex    sync.Mutex
	packages map[types.PackageID]*Package // guarded by mutex
}

func NewBuilder(workspace *Workspace) *Builder {
	ctx, cancel := context.WithCancel(context.Background())
	return &Builder{
		ctx:       ctx,
		cancel:    cancel,
		Emitter:   &errors.Emitter{},
		Workspace: workspace,
		packages:  map[types.PackageID]*Package{},
	}
}

// TODO rm
func (builder *Builder) Packages() map[types.PackageID]*Package {
	builder.buildWaitGroup.Wait()

	builder.mutex.Lock()
	defer builder.mutex.Unlock()

	return builder.packages
}

func (builder *Builder) Build(ids ...types.PackageID) []error {
	for _, id := range ids {
		builder.build(id, nil)
	}

	builder.loadWaitGroup.Wait()

	builder.detectCycles()

	builder.buildWaitGroup.Wait()

	errs := builder.Errors()
	for _, pkg := range builder.packages {
		errs = append(errs, pkg.Errors()...)
	}

	return errs
}

func (builder *Builder) build(
	id types.PackageID,
	by *importedBy,
) *Package {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()

	pkg := builder.packages[id]
	if pkg == nil {
		emitter := &errors.Emitter{}
		pkg = &Package{
			builder:   builder,
			PackageID: id,
			Emitter:   emitter,
			locateState: locateState{
				Emitter: emitter,
			},
			DirectDependencies:      map[types.PackageID]depPkg{},
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

func (builder *Builder) detectCycles() {
	for _, pkg := range builder.packages {
		if pkg.isAcyclic {
			continue
		}

		// NOTE: There could be multiple import cycles, but for simplicity, we'll
		// only report the first detected cycle.
		err := pkg.detectCycle(nil)
		if err != nil {
			builder.EmitErrors(err)
			builder.cancel()
			return
		}
	}
}
