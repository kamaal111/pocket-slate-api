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
		JSON().
		Object()
	makeTranslation.Value("message").IsEqual("Invalid payload provided")
}
