package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/html"
)

func TestIndexRoute(t *testing.T) {
	testService := test.NewService()
	resultsService := results.NewService()

	r := NewRouter(testService, resultsService)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", http.NoBody)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	_, err := html.Parse(w.Body)
	assert.NoError(t, err)
}

func TestEditorUpload(t *testing.T) {
	testService := test.NewService()
	resultsService := results.NewService()

	r := NewRouter(testService, resultsService)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/editor/upload", http.NoBody)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	_, err := html.Parse(w.Body)
	assert.NoError(t, err)
}

func TestEditorEdit(t *testing.T) {
	testService := test.NewService()
	resultsService := results.NewService()

	r := NewRouter(testService, resultsService)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/editor/edit", http.NoBody)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	_, err := html.Parse(w.Body)
	assert.NoError(t, err)
}
