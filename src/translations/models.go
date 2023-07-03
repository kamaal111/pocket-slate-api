package translations

type supportedLocale struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type makeTranslationPayload struct {
	Text         *string `json:"text"`
	TargetLocale *string `json:"target_locale"`
	SourceLocale *string `json:"source_locale"`
}

type makeTranslationResponse struct {
	TranslatedText string `json:"translated_text"`
}
