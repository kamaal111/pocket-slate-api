package health_test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"

	"github.com/kamaal111/pocket-slate-api/src/health"
)

func TestPing(t *testing.T) {
	engine := gin.New()

	handler := health.Router(engine)

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewCookieJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	ping := e.GET("/v1/health/ping").
		Expect().
		Status(http.StatusOK).
		ContentType("application/json").
		JSON().
		Object()
	ping.Value("message").IsEqual("pong")
}
