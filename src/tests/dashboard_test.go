package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
	"github.com/Real-Dev-Squad/gopher-cloud-service/src/tests/fixtures"
	"github.com/stretchr/testify/assert"
)

var endpoint string

func TestHealthDashboardUnauthenticated(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", endpoint, nil)

	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.JSONEq(t, fixtures.Unauthorized(), w.Body.String())
}

func TestHealthDashboardAuthenticated(t *testing.T) {

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", endpoint, nil)

	req.AddCookie(&http.Cookie{
		Name:  config.Global.CookieName,
		Value: fixtures.SendCookie(),
	})

	Router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, fixtures.Authorized(), w.Body.String())
}

func init() {
	TestSetup()
	endpoint = Ver + "/health/dashboard"
}
