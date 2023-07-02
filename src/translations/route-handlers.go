package translations

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"cloud.google.com/go/translate"
	"github.com/kamaal111/pocket-slate-api/src/utils"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func getSupportedLocalesHandler(writer http.ResponseWriter, request *http.Request) {
	translateAPIKey := utils.UnwrapEnvironment("TRANSLATE_API_KEY")
	target := request.URL.Query().Get("target")
	if target == "" {
		utils.ErrorHandler(writer, "target is expected in the query params", http.StatusBadRequest)
		return
	}

	lang, err := language.Parse(target)
	if err != nil {
		utils.ErrorHandler(writer, fmt.Sprintf("Invalid target of %s", target), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	client, err := translate.NewClient(ctx, option.WithAPIKey(translateAPIKey))
	if err != nil {
		utils.ErrorHandler(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer client.Close()

	resp, err := client.SupportedLanguages(ctx, lang)
	if err != nil {
		utils.ErrorHandler(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(resp) == 0 {
		utils.ErrorHandler(writer, "No supported locales found", http.StatusNotFound)
		return
	}

	mappedResponse := utils.MapSlice(resp, func(item translate.Language) supportedLocale {
		return supportedLocale{Name: item.Name, Tag: item.Tag.String()}
	})
	output, err := json.Marshal(mappedResponse)
	if err != nil {
		utils.ErrorHandler(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("content-type", "application/json")
	writer.Write(output)
}
