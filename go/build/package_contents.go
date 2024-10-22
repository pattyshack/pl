package build

import (
	"bytes"
	"crypto/md5"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

const (
	sourceExtension = ".pl"
)

type SourceType int

const (
	NotSource     = SourceType(0)
	LibrarySource = SourceType(1)
	TestSource    = SourceType(2)
)

func IsSource(name string) bool {
	return filepath.Ext(name) == sourceExtension
}

func ClassifySource(fileName string) SourceType {
	if !IsSource(fileName) {
		return NotSource
	}

	fileName = filepath.Base(fileName[:len(fileName)-len(sourceExtension)])
	if strings.HasSuffix(fileName, "_test") ||
		strings.HasSuffix(fileName, "-test") ||
		strings.HasSuffix(fileName, "Test") ||
		strings.HasPrefix(fileName, "test") {

		return TestSource
	}

	return LibrarySource
}

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

	mutex            sync.Mutex
	hasCachedContent bool   // guarded by mutex
	cachedContent    []byte // guarded by mutex
	checksum         []byte // guarded by mutex
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

		if len(file.checksum) == 0 {
			hash := md5.Sum(file.cachedContent)
			file.checksum = hash[:]
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

type PackageContents map[string]File

func (pkg PackageContents) Sources() ([]string, []string) {
	lib := []string{}
	test := []string{}
	for path, file := range pkg {
		switch ClassifySource(file.Name()) {
		case LibrarySource:
			lib = append(lib, path)
		case TestSource:
			test = append(test, path)
		}
	}

	return lib, test
}

func (pkg PackageContents) Open(path string) (io.Reader, error) {
	file, ok := pkg[path]
	if !ok {
		return nil, os.ErrNotExist
	}

	return file.Open()
}

func (pkg PackageContents) ClearCached() {
	for _, file := range pkg {
		file.ClearCached()
	}
}
