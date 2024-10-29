package types

import (
	"fmt"
	"sync"

	"github.com/pattyshack/pl/ast"
)

// TODO: improve data structure
type PropertiesType struct {
	ast.PropertiesKind
	GenericParameters []*ast.GenericParameter

	// All named/unnamed fields/methods
	Fields  []*ast.FieldDef
	Methods []*ast.FuncSignature

	// Unambiguously named fields / methods.  Does not include unnamed fields
	Named map[string]ast.TypeProperty
}

type PackageInterface struct {
	PropertiesType
	TypeDefs map[string]PropertiesType
}

type Catalog struct {
	mutex    sync.Mutex
	packages map[ast.PackageID]PackageInterface // guarded by mutex
}

func NewCatalog() *Catalog {
	return &Catalog{
		packages: map[ast.PackageID]PackageInterface{},
	}
}

func (catalog *Catalog) SetPackageInterface(
	id ast.PackageID,
	pkg PackageInterface,
) {
	catalog.mutex.Lock()
	defer catalog.mutex.Unlock()

	_, ok := catalog.packages[id]
	if ok {
		panic(fmt.Sprintf("setting duplicate package %s", id))
	}

	catalog.packages[id] = pkg
}

func (catalog *Catalog) GetPackageInterface(id ast.PackageID) PackageInterface {
	catalog.mutex.Lock()
	defer catalog.mutex.Unlock()

	return catalog.packages[id]
}
