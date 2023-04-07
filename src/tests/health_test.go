package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/tests/fixtures"

	"github.com/stretchr/testify/assert"
)

func TestHealthRoute(t *testing.T) {

	TestSetup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", Ver+"/health", nil)
	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, fixtures.Health(), w.Body.String())
}
