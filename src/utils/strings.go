package utils

import (
	"strings"
	"unicode"
)

func PascalToSnakeCase(term string) string {
	var words []string
	var word string
	for _, character := range term {
		if unicode.IsUpper(character) {
			if len(word) != 0 {
				words = append(words, word)
				word = ""
			}
		}

		word += string(unicode.ToLower(character))
	}

	if len(word) != 0 {
		words = append(words, word)
	}

	return strings.Join(words, "_")
}
