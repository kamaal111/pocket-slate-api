package translations_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/kamaal111/pocket-slate-api/src/translations"
)

const (
	APP_NAME    = "pocket-slate"
	API_KEY     = "96c64d4a0d615e8a1ec5622acc0589cb004a7948894ccee40c1beb607ab66b2d"
	APP_VERSION = "1.0.0"
)

var HEADERS = map[string]string{
	"Api-Key":     API_KEY,
	"App-Name":    APP_NAME,
	"App-Version": APP_VERSION,
}

func TestMakeTranslationHandlerNoPayloadProvided(t *testing.T) {
	t.Setenv("APP_API_KEYS", fmt.Sprintf("{\"%s\":{\"%s\":\"%s\"}}", APP_NAME, APP_VERSION, API_KEY))

	engine := gin.New()

	handler := translations.Router(engine)

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

	makeTranslation := e.POST("/v1/translations").
		WithHeaders(HEADERS).
		Expect().
		Status(http.StatusBadRequest).
		ContentType("application/json").
		JSON().
		Object()
	makeTranslation.Value("message").IsEqual("Invalid payload provided")
}

func TestMakeTranslationHandlerInCompletePayload(t *testing.T) {
	t.Setenv("APP_API_KEYS", fmt.Sprintf("{\"%s\":{\"%s\":\"%s\"}}", APP_NAME, APP_VERSION, API_KEY))

	payloads := []map[string]interface{}{
		{
			"text":          "Yes",
			"target_locale": "it",
		},
		{
			"text":          "Yes",
			"source_locale": "en",
		},
		{},
		{
			"text": "Yes",
		},
		{
			"source_locale": "en",
		},
		{
			"target_locale": "it",
		},
	}

	engine := gin.New()

	handler := translations.Router(engine)

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

	for _, payload := range payloads {
		makeTranslation := e.POST("/v1/translations").
			WithHeaders(HEADERS).
			WithJSON(payload).
			Expect().
			Status(http.StatusUnprocessableEntity).
			ContentType("application/json").
			JSON().
			Object()
		makeTranslation.Value("message").String().Contains("' is required in the body")
	}

}

func TestGetSupportedLocalesMissingTarget(t *testing.T) {
	t.Setenv("APP_API_KEYS", fmt.Sprintf("{\"%s\":{\"%s\":\"%s\"}}", APP_NAME, APP_VERSION, API_KEY))

	engine := gin.New()

	handler := translations.Router(engine)

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

	supportedLocales := e.GET("/v1/translations/supported-locales").
		WithHeaders(HEADERS).
		Expect().
		Status(http.StatusUnprocessableEntity).
		ContentType("application/json").
		JSON().
		Object()
	supportedLocales.Value("message").IsEqual("'target' is required in the query params")
}
