package build

import (
	"bytes"
	"crypto/md5"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/pattyshack/pl/ast"
)

type File interface {
	Name() string
	Size() int64
	ModTime() time.Time

	Checksum() ([]byte, error)
	Open() (io.Reader, error)

	ClearCached()
}

type localFile struct {
	dir     string
	name    string
	size    int64
	modTime time.Time

	mutex               sync.Mutex
	hasCachedContent    bool   // guarded by mutex
	cachedContent       []byte // guarded by mutex
	initializedChecksum bool   // guarded by mutex
	checksum            []byte // guarded by mutex
}

func (file *localFile) Name() string {
	return file.name
}

func (file *localFile) Size() int64 {
	return file.size
}

func (file *localFile) ModTime() time.Time {
	return file.modTime
}

func (file *localFile) unsafeReadContent() error {
	if !file.hasCachedContent {
		content, err := os.ReadFile(filepath.Join(file.dir, file.name))
		if err != nil {
			return err
		}

		file.cachedContent = content
		file.hasCachedContent = true

		if !file.initializedChecksum {
			hash := md5.Sum(file.cachedContent)
			file.checksum = hash[:]
			file.initializedChecksum = true
		}
	}

	return nil
}

func (file *localFile) Checksum() ([]byte, error) {
	file.mutex.Lock()
	defer file.mutex.Unlock()

	err := file.unsafeReadContent()
	if err != nil {
		return nil, err
	}

	return file.checksum, nil
}

func (file *localFile) Open() (io.Reader, error) {
	file.mutex.Lock()
	defer file.mutex.Unlock()

	err := file.unsafeReadContent()
	if err != nil {
		return nil, err
	}

	content := file.cachedContent
	file.cachedContent = nil
	file.hasCachedContent = false

	return bytes.NewBuffer(content), nil
}

func (file *localFile) ClearCached() {
	file.mutex.Lock()
	file.mutex.Unlock()

	file.hasCachedContent = false
	file.cachedContent = nil
}

type PackageContents struct {
	ast.PackageID

	Contents map[string]File
}

func (pkg PackageContents) Sources() []string {
	sources := []string{}
	for name, _ := range pkg.Contents {
		sources = append(sources, name)
	}

	return sources
}

func (pkg PackageContents) Open(name string) (io.Reader, error) {
	file, ok := pkg.Contents[name]
	if !ok {
		return nil, os.ErrNotExist
	}

	return file.Open()
}
