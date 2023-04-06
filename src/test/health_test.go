package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/routes"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/test/fixtures"

	"github.com/stretchr/testify/assert"
)

func TestHealthRoute(t *testing.T) {
	os.Setenv("env", "test")
	env := os.Getenv("env")
	version := strconv.Itoa(1)
	ver := "/v" + version
	router := routes.SetupRouter(env, version)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", ver+"/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, fixtures.Health(), w.Body.String())
}
