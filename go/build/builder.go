package build

import (
	"context"
	"sync"

	"github.com/pattyshack/gt/lexutil"

	"github.com/pattyshack/pl/ast"
	"github.com/pattyshack/pl/parser"
)

type Package struct {
	builder   *Builder
	importLoc *lexutil.Location

	ast.PackageID
	PackageContents

	HasErrors   bool // maybe classify error type: body error vs public api error
	Definitions *ast.StatementList

	DirectDependencies      map[ast.PackageID]*Package
	PublicInterfaceAnalyzed chan struct{}
}

func (pkg *Package) Load() {
	defer pkg.builder.loadWaitGroup.Done()

	// TODO load cached entry

	contents, err := pkg.builder.Locate(pkg.PackageID)
	if err != nil {
		if pkg.importLoc != nil {
			err = lexutil.LocationError{Loc: *pkg.importLoc, Err: err}
		}
		pkg.builder.EmitErrors(err)
		pkg.HasErrors = true
		return
	}

	parseOptions := parser.ParserOptions{
		OpenFunc: contents.Open,
	}

	emitter := &lexutil.ErrorEmitter{}
	definitions, importPkgs := parser.ParsePackage(
		contents.Sources(),
		emitter,
		parseOptions)
	pkg.Definitions = definitions
	pkg.HasErrors = emitter.HasErrors()
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

func (builder *Builder) Build(ids ...ast.PackageID) {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()

	for _, id := range ids {
		builder.unsafeBuild(id, nil)
	}
}

func (builder *Builder) build(
	id ast.PackageID,
	importLoc *lexutil.Location,
) *Package {
	builder.mutex.Lock()
	defer builder.mutex.Unlock()

	return builder.unsafeBuild(id, importLoc)
}

func (builder *Builder) unsafeBuild(
	id ast.PackageID,
	importLoc *lexutil.Location,
) *Package {
	pkg := builder.packages[id]
	if pkg != nil {
		return pkg
	}

	analyzedChan := make(chan struct{})

	pkg = &Package{
		builder:                 builder,
		importLoc:               importLoc,
		PackageID:               id,
		DirectDependencies:      map[ast.PackageID]*Package{},
		PublicInterfaceAnalyzed: analyzedChan,
	}
	builder.packages[id] = pkg

	builder.loadWaitGroup.Add(1)
	builder.buildWaitGroup.Add(1)
	go pkg.Build()

	return pkg
}

func (builder *Builder) WaitForCompletion() []error {
	builder.loadWaitGroup.Wait()

	// TODO check for dep cycles, cancel if found cycle

	builder.buildWaitGroup.Wait()
	return builder.Errors()
}
