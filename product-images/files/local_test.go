package files

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupLocal(t *testing.T) (*Local, string, func()) {
	dir, err := ioutil.TempDir(".", "files")
	if err != nil {
		t.Fatal(err)
	}

	l, err := NewLocal(dir, -1)
	if err != nil {
		t.Fatal(err)
	}
	return l, dir, func() {
		// cleanup function
		os.RemoveAll(dir)
	}
}

func TestSaveContentsOfReader(t *testing.T) {
	savePath := "/1/test.png"
	fileContents := "Hello World"
	l, dir, cleanUp := setupLocal(t)
	defer cleanUp()

	err := l.Save(savePath, bytes.NewBuffer([]byte(fileContents)))
	assert.NoError(t, err)

	f, err := os.Open(filepath.Join(dir, savePath))
	assert.NoError(t, err)

	d, err := ioutil.ReadAll(f)
	assert.NoError(t, err)
	assert.Equal(t, fileContents, string(d))
}

func TestGetContentsAndWritesToWriter(t *testing.T) {
	savepath := "/1/test.png"
	fileContents := "Hello World"
	l, _, cleanUp := setupLocal(t)
	defer cleanUp()

	err := l.Save(savepath, bytes.NewBuffer([]byte(fileContents)))
	assert.NoError(t, err)

	r, err := l.Get(savepath)
	assert.NoError(t, err)
	defer r.Close()

	d, err := ioutil.ReadAll(r)
	assert.NoError(t, err)
	assert.Equal(t, fileContents, string(d))
}
