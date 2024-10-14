package build

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pattyshack/pl/ast"
)

const (
	workspaceFile   = "pl.workspace"
	sourceExtension = ".pl"
)

// TODO: support modularization.  Map all packages, including external modules
// relative to // root (don't support bazel external repo style)
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

func (workspace *Workspace) Locate(id ast.PackageID) (PackageContents, error) {
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

	contents := PackageContents{}
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

	return contents, nil
}

// This uses bazel's cmdline convention (with no external repo) rather than
// go's cmdline convention. In particular, // is the fully qualified package
// path prefix (always uses slash regardless of filesystem separator).  ... by
// itself is relative to cwd rather than //.
//
// (note that workspace root may or may not be the same as // once
// modularization is supported).
//
// TODO: modularization
func (workspace *Workspace) MatchTargetPattern(
	originalPattern string,
) (
	[]ast.PackageID,
	error,
) {
	pattern := originalPattern
	pkgRoot := ""
	findSubPackages := false

	if pattern == "..." {
		pkgRoot = workspace.WorkingDirPath
		findSubPackages = true
	} else if strings.HasPrefix(pattern, "//") {
		pattern = pattern[2:]
		if pattern == "" {
			pkgRoot = workspace.RootDirPath
		} else if pattern == "..." {
			pkgRoot = workspace.RootDirPath
			findSubPackages = true
		} else {
			if strings.HasSuffix(pattern, "/...") {
				findSubPackages = true
				pattern = pattern[:len(pattern)-4]
			}

			cleaned := path.Clean(pattern)
			if cleaned == "." || cleaned != pattern {
				return nil, fmt.Errorf(
					"invalid pattern (%s), pattern package path is not clean/absolute",
					originalPattern)
			}

			pkgRoot = filepath.Join(
				workspace.RootDirPath,
				filepath.FromSlash(pattern))
		}
	} else {
		dir, file := filepath.Split(pattern)
		if file == "..." {
			findSubPackages = true
			pattern = dir
		}

		var err error
		pkgRoot, err = filepath.Abs(pattern)
		if err != nil {
			return nil, fmt.Errorf("invalid pattern (%s): %w", originalPattern, err)
		}

		if pkgRoot != workspace.RootDirPath &&
			!strings.HasPrefix(
				pkgRoot,
				workspace.RootDirPath+string(filepath.Separator)) {

			return nil, fmt.Errorf(
				"invalid pattern (%s), not in workspace (%s)",
				originalPattern,
				workspace.RootDirPath)
		}
	}

	if !findSubPackages && pkgRoot == workspace.RootDirPath {
		return nil, fmt.Errorf(
			"invalid pattern path (%s), must specify package name (or '...')",
			originalPattern)
	}

	info, err := os.Stat(pkgRoot)
	if err != nil {
		return nil, fmt.Errorf(
			"invalid pattern path (%s): %w",
			originalPattern,
			err)
	} else if !info.IsDir() {
		return nil, fmt.Errorf(
			"invalid pattern (%s), pattern package path is not a directory",
			originalPattern)
	}

	dirs := map[string]struct{}{}
	if !findSubPackages {
		dirs[pkgRoot] = struct{}{}
	} else {
		err := filepath.WalkDir(
			pkgRoot,
			func(path string, entry fs.DirEntry, err error) error {
				if entry.IsDir() || filepath.Ext(entry.Name()) != sourceExtension {
					return nil
				}

				dir, _ := filepath.Split(path)
				dirs[dir] = struct{}{}
				return nil
			})
		if err != nil {
			return nil, fmt.Errorf("invalid pattern (%s): %w", originalPattern, err)
		}
	}

	targets := []ast.PackageID{}
	for dir, _ := range dirs {
		relDir, err := filepath.Rel(workspace.RootDirPath, dir)
		if err != nil {
			return nil, fmt.Errorf("invalid pattern (%s): %w", originalPattern, err)
		}

		id := ast.NewPackageID(filepath.ToSlash(relDir))
		err = id.Validate()
		if err != nil {
			return nil, fmt.Errorf("found invalid package %s: %w", dir, err)
		} else {
			targets = append(targets, id)
		}
	}

	return targets, nil
}

func (workspace *Workspace) MatchTargetPatterns(
	patterns ...string,
) (
	[]ast.PackageID,
	error,
) {
	pkgIds := []ast.PackageID{}
	for _, pattern := range patterns {
		ids, err := workspace.MatchTargetPattern(pattern)
		if err != nil {
			return nil, err
		}

		pkgIds = append(pkgIds, ids...)
	}

	return pkgIds, nil
}
