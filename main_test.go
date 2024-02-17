package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	jsonFile, err := os.ReadFile("init-example.json")
	assert.Nil(t, err)
	req, _ := http.NewRequest("POST", "/generate", bytes.NewBuffer(jsonFile))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, "application/zip", w.Header().Get("Content-Type"))
	assert.Equal(t, 200, w.Code)
}
