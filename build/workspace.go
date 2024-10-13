package build

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pattyshack/pl/ast"
)

const (
	workspaceFile   = "pl.workspace"
	sourceExtension = ".pl"
)

type Workspace struct {
	RootDirPath    string
	WorkingDirPath string
}

func FindWorkspaceRoot() (*Workspace, error) {
	dir, err := filepath.Abs(".")
	if err != nil {
		return nil, fmt.Errorf("unable to locate workspace directory: %w", err)
	}

	workingDir := dir

	for {
		entries, err := os.ReadDir(dir)
		if err != nil {
			return nil, fmt.Errorf("unable to locate workspace directory: %w", err)
		}

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			if entry.Name() == workspaceFile {
				return &Workspace{
					RootDirPath:    dir,
					WorkingDirPath: workingDir,
				}, nil
			}
		}

		parentDir := filepath.Dir(dir)
		if dir == parentDir {
			return nil, fmt.Errorf(
				"unable to locate workspace directory, %s file not found",
				workspaceFile)
		}
		dir = parentDir
	}
}

func (workspace *Workspace) Locate(id ast.PackageID) (*PackageContents, error) {
	err := id.Validate()
	if err != nil {
		return nil, err
	}

	pkgDirPath := filepath.Join(
		workspace.RootDirPath,
		filepath.FromSlash(id.String()))

	entries, err := os.ReadDir(pkgDirPath)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to locate package (%s): %w",
			id.String(),
			err)
	}

	contents := map[string]File{}
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != sourceExtension {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			return nil, fmt.Errorf(
				"failed to locate package (%s): %w",
				id.String(),
				err)
		}

		relPath, err := filepath.Rel(
			workspace.WorkingDirPath,
			filepath.Join(pkgDirPath, info.Name()))
		if err != nil {
			return nil, fmt.Errorf(
				"failed to locate package (%s): %w",
				id.String(),
				err)
		}

		contents[relPath] = &localFile{
			dir:     pkgDirPath,
			name:    info.Name(),
			size:    info.Size(),
			modTime: info.ModTime(),
		}
	}

	if len(contents) == 0 {
		return nil, fmt.Errorf(
			"invalid package (%s), package directory does not contain source files",
			id.String())
	}

	return &PackageContents{
		PackageID: id,
		Contents:  contents,
	}, nil
}
