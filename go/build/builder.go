package build

import (
	"context"
	"fmt"
	"sync"

	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser"
)

type locateState struct {
	builder *Builder

	ast.PackageID

	mutex sync.Mutex

	// All guarded by mutex
	located       bool
	err           error
	isBuildTarget bool
	importedBy    []lexutil.Location
}

func (state *locateState) Locate() (
	PackageContents,
	error,
) {
	state.mutex.Lock()
	defer state.mutex.Unlock()

	if state.located {
		panic("should never happen")
	}
	state.located = true

	contents, err := state.builder.Locate(state.PackageID)
	state.err = err

	if err != nil {
		if state.isBuildTarget {
			state.builder.EmitErrors(err)
		}

		for _, loc := range state.importedBy {
			state.builder.EmitErrors(lexutil.LocationError{Loc: loc, Err: err})
		}
	}
	state.importedBy = nil

	return contents, err
}

func (state *locateState) ImportedBy(loc *lexutil.Location) {
	state.mutex.Lock()
	defer state.mutex.Unlock()

	if !state.located {
		if loc == nil {
			state.isBuildTarget = true
		} else {
			state.importedBy = append(state.importedBy, *loc)
		}
		return
	}

	if loc == nil {
		if state.isBuildTarget {
			return // already emitted error, if any
		}
		state.isBuildTarget = true

		if state.err != nil {
			state.builder.EmitErrors(state.err)
		}
	} else if state.err != nil {
		state.builder.EmitErrors(lexutil.LocationError{Loc: *loc, Err: state.err})
	}
}

type Package struct {
	builder *Builder
	locateState

	// Populated by Load().  Safe to access once loadWaitGroup is ready
	HasLoadErrors bool
	PackageContents
	Definitions        *ast.StatementList
	DirectDependencies map[ast.PackageID]*Package

	// Only used by detectCycles, after loadWaitGroup is ready.
	isAcyclic bool

	PublicInterfaceAnalyzed chan struct{}
}

func (pkg *Package) Load() {
	defer pkg.builder.loadWaitGroup.Done()

	// TODO load cached entry

	contents, err := pkg.locateState.Locate()
	if err != nil {
		pkg.HasLoadErrors = true
		return
	}
	pkg.PackageContents = contents

	parseOptions := parser.ParserOptions{
		OpenFunc: contents.Open,
	}

	emitter := &lexutil.ErrorEmitter{}
	definitions, importPkgs := parser.ParsePackage(
		contents.Sources(),
		emitter,
		parseOptions)
	pkg.Definitions = definitions
	pkg.HasLoadErrors = emitter.HasErrors()
	pkg.builder.ErrorEmitter.MergeFrom(emitter)

	for _, importPkg := range importPkgs {
		loc := importPkg.Loc()
		pkg.DirectDependencies[importPkg.PackageID] = pkg.builder.build(
			importPkg.PackageID,
			&loc)
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

func (pkg *Package) detectCycle(visited []ast.PackageID) error {
	for idx, visitedID := range visited {
		if visitedID == pkg.PackageID {
			cycleErrMsg := "detected import dependency cycle:\n  "
			for _, id := range visited[idx:] {
				cycleErrMsg += id.String() + " ->\n  "
			}
			cycleErrMsg += pkg.PackageID.String()

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

	*lexutil.ErrorEmitter

	loadWaitGroup  sync.WaitGroup
	buildWaitGroup sync.WaitGroup

	*Workspace

	mutex    sync.Mutex
	packages map[ast.PackageID]*Package // guarded by mutex
}

func NewBuilder(workspace *Workspace) *Builder {
	ctx, cancel := context.WithCancel(context.Background())
	return &Builder{
		ctx:          ctx,
		cancel:       cancel,
		ErrorEmitter: &lexutil.ErrorEmitter{},
		Workspace:    workspace,
		packages:     map[ast.PackageID]*Package{},
	}
}

// TODO rm
func (builder *Builder) Packages() map[ast.PackageID]*Package {
	builder.buildWaitGroup.Wait()

	builder.mutex.Lock()
	defer builder.mutex.Unlock()

	return builder.packages
}

func (builder *Builder) Build(ids ...ast.PackageID) []error {
	for _, id := range ids {
		builder.build(id, nil)
	}

	builder.loadWaitGroup.Wait()

	builder.detectCycles()

	builder.buildWaitGroup.Wait()
	return builder.Errors()
}

func (builder *Builder) build(
	id ast.PackageID,
	importLoc *lexutil.Location,
) *Package {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()

	pkg := builder.packages[id]
	if pkg == nil {
		pkg = &Package{
			builder: builder,
			locateState: locateState{
				builder:   builder,
				PackageID: id,
			},
			DirectDependencies:      map[ast.PackageID]*Package{},
			PublicInterfaceAnalyzed: make(chan struct{}),
		}
		builder.packages[id] = pkg

		builder.loadWaitGroup.Add(1)
		builder.buildWaitGroup.Add(1)
		go pkg.Build()
	}

	pkg.locateState.ImportedBy(importLoc)
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
