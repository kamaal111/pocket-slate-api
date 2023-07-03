package translations

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/translate"
	"github.com/kamaal111/pocket-slate-api/src/utils"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

func withTranslationService(callback func(translationService)) error {
	translateAPIKey, err := utils.UnwrapEnvironment("TRANSLATE_API_KEY")
	if err != nil {
		return err
	}

	ctx := context.Background()
	client, err := translate.NewClient(ctx, option.WithAPIKey(translateAPIKey))
	if err != nil {
		return err
	}
	defer client.Close()

	service := translationService{Client: client, Context: ctx}
	callback(service)

	return nil
}

type translationService struct {
	Client  *translate.Client
	Context context.Context
}

func (service *translationService) Translate(text string, source string, target string) (string, *utils.Error) {
	targetLang, err := language.Parse(target)
	if err != nil {
		return "", &utils.Error{
			Message: fmt.Sprintf("Invalid target of %s", target),
			Status:  http.StatusBadRequest,
		}
	}

	sourceLang, err := language.Parse(source)
	if err != nil {
		return "", &utils.Error{
			Message: fmt.Sprintf("Invalid source of %s", target),
			Status:  http.StatusBadRequest,
		}
	}

	resp, err := service.Client.Translate(service.Context, []string{text}, targetLang, &translate.Options{Source: sourceLang})
	if err != nil {
		return "", &utils.Error{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if len(resp) == 0 {
		return "", &utils.Error{
			Message: "No translations found",
			Status:  http.StatusNotFound,
		}
	}

	return resp[0].Text, nil
}

func (service *translationService) SupportedLanguages(target string) ([]supportedLocale, *utils.Error) {
	lang, err := language.Parse(target)
	if err != nil {
		return []supportedLocale{}, &utils.Error{
			Message: fmt.Sprintf("Invalid target of %s", target),
			Status:  http.StatusBadRequest,
		}
	}

	resp, err := service.Client.SupportedLanguages(service.Context, lang)
	if err != nil {
		return []supportedLocale{}, &utils.Error{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if len(resp) == 0 {
		return []supportedLocale{}, &utils.Error{
			Message: "No supported locales found",
			Status:  http.StatusNotFound,
		}
	}

	mappedResponse := utils.MapSlice(resp, func(item translate.Language) supportedLocale {
		return supportedLocale{Name: item.Name, Tag: item.Tag.String()}
	})
	return mappedResponse, nil
}
