package routers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestErrors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(notFound))
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	notFoundResponse := e.GET("/").
		Expect().
		Status(http.StatusNotFound).
		JSON().
		Object()
	notFoundResponse.Value("message").IsEqual("Not found")
}
