package handlers

import (
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/terrytay/microservices-with-go/product-images/files"
)

type Files struct {
	log      *log.Logger
	store    files.Storage
	urlParam func(r *http.Request, key string) string
}

// NewFiles takes in a logger, a storage and a URLparams function
func NewFiles(l *log.Logger, s files.Storage, m func(r *http.Request, key string) string) *Files {
	return &Files{log: l, store: s, urlParam: m}
}

func (f *Files) UploadREST(rw http.ResponseWriter, r *http.Request) {
	id := f.urlParam(r, "id")
	fn := f.urlParam(r, "filename")

	f.log.Println("[DEBUG] Handle POST", "id", id, "filename", fn)

	f.saveFile(id, fn, rw, r.Body)
}

func (f *Files) UploadMultipart(rw http.ResponseWriter, r *http.Request) {
	f.log.Println("[DEBUG] Handle POST")

	err := r.ParseMultipartForm(128 * 1024) // 128kB
	if err != nil {
		f.log.Println("[DEBUG] Bad Request", "error", err)
		http.Error(rw, "expected multipart form data", http.StatusBadRequest)
		return
	}

	id, idErr := strconv.Atoi(r.FormValue("id"))
	f.log.Println("[DEBUG] Process form for id", id)
	if idErr != nil {
		f.log.Println("[DEBUG] Bad Request", "error", err)
		http.Error(rw, "expected integer id", http.StatusBadRequest)
		return
	}

	ff, mh, err := r.FormFile("file")
	if err != nil {
		f.log.Println("[DEBUG] Bad Request", "error", err)
		http.Error(rw, "expected file", http.StatusBadRequest)
		return
	}

	f.saveFile(r.FormValue("id"), mh.Filename, rw, ff)
}

func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r io.ReadCloser) {
	f.log.Println("[DEBUG] Save file")

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r)
	if err != nil {
		f.log.Println("[DEBUG] Unable to save file")
		http.Error(rw, "unaable to save file", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
