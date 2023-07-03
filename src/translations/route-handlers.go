package translations

import (
	"encoding/json"
	"net/http"

	"github.com/kamaal111/pocket-slate-api/src/utils"
)

func makeTranslationHandler(writer http.ResponseWriter, request *http.Request) {
	var payload makeTranslationPayload
	err := json.NewDecoder(request.Body).Decode(&payload)
	if err != nil {
		utils.ErrorHandler(writer, "Invalid payload provided", http.StatusBadRequest)
		return
	}

	if payload.Text == nil || payload.SourceLocale == nil || payload.TargetLocale == nil {
		utils.ErrorHandler(writer, "Invalid payload provided", http.StatusBadRequest)
		return
	}

	var resp string
	var httpErr *utils.Error
	err = withTranslationService(func(ts translationService) {
		resp, httpErr = ts.Translate(*payload.Text, *payload.SourceLocale, *payload.TargetLocale)
	})
	if err != nil {
		utils.ErrorHandler(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if httpErr != nil {
		utils.ErrorHandler(writer, httpErr.Message, httpErr.Status)
		return
	}

	utils.MarshalJSONResponse(writer, makeTranslationResponse{TranslatedText: resp})
}

func getSupportedLocalesHandler(writer http.ResponseWriter, request *http.Request) {
	target, err := utils.UnwrapURLQuery(request, "target")
	if target == "" {
		utils.ErrorHandler(writer, err.Error(), http.StatusBadRequest)
		return
	}

	var resp []supportedLocale
	var httpErr *utils.Error
	err = withTranslationService(func(ts translationService) {
		resp, httpErr = ts.SupportedLanguages(target)
	})
	if err != nil {
		utils.ErrorHandler(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	if httpErr != nil {
		utils.ErrorHandler(writer, httpErr.Message, httpErr.Status)
		return
	}

	utils.MarshalJSONResponse(writer, resp)
}
