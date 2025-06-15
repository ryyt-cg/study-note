package main

import (
	"github.com/json-iterator/go"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// How to write test case for Gin?
// The net/http/httptest package is a preferable way for HTTP testing.
// It provides a ResponseRecorder which is an implementation of http.ResponseWriter
// that records its mutations for later inspection in tests.
func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

// Test for POST /user/add
func TestPostUser(t *testing.T) {
	router := setupRouter()
	router = postUser(router)

	w := httptest.NewRecorder()
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	// Create an example user for testing
	exampleUser := User{
		Username: "test_name",
		Gender:   "male",
	}
	userJson, _ := json.Marshal(exampleUser)
	req, _ := http.NewRequest("POST", "/user/add", strings.NewReader(string(userJson)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// Compare the response body with the json data of exampleUser
	assert.Equal(t, string(userJson), w.Body.String())
}
