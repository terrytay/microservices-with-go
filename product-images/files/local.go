package files

import (
	"io"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

type Local struct {
	maxFileSize int // max bytes for files
	basePath    string
}

// NewLocal creates the basePath by getting the absolute path and store in Local struct
func NewLocal(basePath string, maxSize int) (*Local, error) {
	p, err := filepath.Abs(basePath)
	if err != nil {
		return nil, err
	}

	// return &Local{basePath: p, maxFileSize: maxSize}, nil
	return &Local{basePath: p}, nil
}

// Save saves the content of the Writer to the given path
// path is a relative path and will be pre-appended by basePath
func (l *Local) Save(path string, contents io.Reader) error {
	// get full path for file
	fp := l.fullPath(path)

	// get the directory and make sure it exists
	d := filepath.Dir(fp)
	err := os.Mkdir(d, os.ModePerm)
	if err != nil {
		return xerrors.Errorf("unable to create directory: %w", err) // %w unwraps the error which makes it accessible to errors.Is
	}

	// if file exists, delete it
	_, err = os.Stat(fp)
	if err == nil {
		err = os.Remove(fp)
		if err != nil {
			return xerrors.Errorf("unable to delete file: %w", err)
		}
	} else if !os.IsNotExist(err) {
		return xerrors.Errorf("unable to get file info: %w", err)
	}

	// creates a new file at the path
	f, err := os.Create(fp)
	if err != nil {
		xerrors.Errorf("unable to create file: %w", err)
	}
	defer f.Close()

	// write the contents to the file ~without exceeding max size~
	_, err = io.Copy(f, contents)
	if err != nil {
		return xerrors.Errorf("unable to write to file: %w", err)
	}

	// return nil if no errors
	return nil
}

// Get returns a reader to the file at the specified path
func (l *Local) Get(path string) (*os.File, error) {
	fp := l.fullPath(path)
	f, err := os.Open(fp)

	if err != nil {
		return nil, err
	}

	return f, nil
}

// fullPath returns full path
func (l *Local) fullPath(path string) string {
	return filepath.Join(l.basePath, path)
}
