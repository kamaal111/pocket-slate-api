package health_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/kamaal111/pocket-slate-api/src/health"
)

func TestPing(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(health.PingHandler))
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	ping := e.GET("/").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()
	ping.Value("message").IsEqual("pong")
}
