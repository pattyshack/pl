package ast

import (
	"fmt"
	"path"
	"strings"

	"github.com/pattyshack/gt/parseutil"
)

type PackageID struct {
	Path    string
	Version string
}

func NewPackageID(pkgVer string) PackageID {
	path, version, _ := strings.Cut(pkgVer, "@")
	return PackageID{
		Path:    path,
		Version: version,
	}
}

func (id PackageID) String() string {
	if id.Version != "" {
		return id.Path + "@" + id.Version
	}

	return id.Path
}

func (id PackageID) Name() string {
	return path.Base(id.Path)
}

func (id PackageID) Validate() error {
	idStr := path.Clean(id.String())
	if strings.HasPrefix(idStr, "/") ||
		strings.HasPrefix(idStr, ".") ||
		strings.HasPrefix(idStr, "@") {

		return fmt.Errorf("invalid pacakge id: %s", idStr)
	}

	if idStr != id.String() {
		return fmt.Errorf("invalid package id, path is not clean: %s", id.String())
	}

	pkgName := id.Name()
	if !parseutil.IsValidIdentifier(pkgName) || pkgName == "_" {
		return fmt.Errorf(
			"invalid package name (%v), package name must be a valid identifier",
			pkgName)
	}

	return nil
}
