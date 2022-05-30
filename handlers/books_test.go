package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/santosh/gingo/models"
	"github.com/santosh/gingo/routes"
	"github.com/stretchr/testify/assert"
)

func TestBooksRoute(t *testing.T) {
	router := routes.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "1")
	assert.Contains(t, w.Body.String(), "2")
	assert.Contains(t, w.Body.String(), "3")
}
func TestBooksbyISBNRoute(t *testing.T) {
	router := routes.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "When Nietzsche Wept")
}
func TestPostBookRoute(t *testing.T) {
	router := routes.SetupRouter()
	book := models.Book{
		ISBN:   "2",
		Author: "Sigmund Freud",
		Title:  "Female Homosexuality",
	}
	body, _ := json.Marshal(book)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/books", bytes.NewReader(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
	assert.Contains(t, w.Body.String(), "Female Homosexuality")
}
