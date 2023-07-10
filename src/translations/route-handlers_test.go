package translations_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/kamaal111/pocket-slate-api/src/translations"
)

func TestMakeTranslationHandlerNoPayloadProvided(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(translations.MakeTranslationHandler))
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	makeTranslation := e.POST("/").
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

	server := httptest.NewServer(http.HandlerFunc(translations.MakeTranslationHandler))
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	for _, payload := range payloads {
		makeTranslation := e.POST("/").
			WithJSON(payload).
			Expect().
			Status(http.StatusUnprocessableEntity).
			ContentType("application/json").
			JSON().
			Object()
		makeTranslation.Value("message").IsEqual("Incomplete payload provided")
	}

}

func TestGetSupportedLocalesMissingTarget(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(translations.GetSupportedLocalesHandler))
	defer server.Close()

	e := httpexpect.Default(t, server.URL)

	makeTranslation := e.GET("/").
		Expect().
		Status(http.StatusBadRequest).
		ContentType("application/json").
		JSON().
		Object()
	makeTranslation.Value("message").IsEqual("'target' is not defined in the query params")
}
