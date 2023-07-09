package routers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"

	"github.com/kamaal111/pocket-slate-api/src/routers"
)

func TestErrors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(routers.NotFound))
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	notFoundResponse := e.GET("/").
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object()
	notFoundResponse.Value("message").IsEqual("Not found")
}
