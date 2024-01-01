package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestIndexRoute(t *testing.T) {
	r := NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", http.NoBody)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	_, err := html.Parse(w.Body)
	assert.NoError(t, err)
}

func TestEditorUpload(t *testing.T) {
	r := NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/editor/upload", http.NoBody)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	_, err := html.Parse(w.Body)
	assert.NoError(t, err)
}

func TestEditorEdit(t *testing.T) {
	r := NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/editor/edit", http.NoBody)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	_, err := html.Parse(w.Body)
	assert.NoError(t, err)
}
