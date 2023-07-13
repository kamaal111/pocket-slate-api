package translations_test

import (
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
	"github.com/kamaal111/pocket-slate-api/src/translations"
)

func TestMakeTranslationHandlerNoPayloadProvided(t *testing.T) {
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
		Expect().
		Status(http.StatusBadRequest).
		ContentType("application/json").
		JSON().
		Object()
	makeTranslation.Value("message").IsEqual("Invalid payload provided")
}

func TestMakeTranslationHandlerInCompletePayload(t *testing.T) {
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
			WithJSON(payload).
			Expect().
			Status(http.StatusUnprocessableEntity).
			ContentType("application/json").
			JSON().
			Object()
		makeTranslation.Value("message").String().Contains("is required")
	}

}

func TestGetSupportedLocalesMissingTarget(t *testing.T) {
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
		Expect().
		Status(http.StatusUnprocessableEntity).
		ContentType("application/json").
		JSON().
		Object()
	supportedLocales.Value("message").IsEqual("'target' is required")
}
