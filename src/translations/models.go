package translations

type supportedLocaleResponse struct {
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type makeTranslationPayload struct {
	Text         string `json:"text" binding:"required"`
	TargetLocale string `json:"target_locale" binding:"required"`
	SourceLocale string `json:"source_locale" binding:"required"`
}

type getSupportedLocalesQuery struct {
	Target string `form:"target"`
}
