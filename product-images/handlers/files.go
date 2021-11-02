package handlers

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/terrytay/microservices-with-go/product-images/files"
)

type Files struct {
	log   *log.Logger
	store files.Storage
}

func NewFiles(l *log.Logger, s files.Storage) *Files {
	return &Files{log: l, store: s}
}

func (f *Files) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fn := chi.URLParam(r, "filename")

	f.log.Println("[DEBUG] Handle POST", "id", id, "filename", fn)

	f.saveFile(id, fn, rw, r)
}

func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r *http.Request) {
	f.log.Println("[DEBUG] Save file")

	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r.Body)
	if err != nil {
		f.log.Println("[DEBUG] Unable to save file")
		http.Error(rw, "unaable to save file", http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
