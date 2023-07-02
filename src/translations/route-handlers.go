package translations

import (
	"net/http"

	"github.com/kamaal111/pocket-slate-api/src/utils"
)

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
