package camelCase

import (
	"strings"
	"unicode"
)

func ToCamelCase(s string) string {
	result := ""

	var split []string
	if s == "" {
		return ""
	}

	if strings.Contains(s, "-") {
		split = strings.Split(s, "-")
	} else if strings.Contains(s, "_") {
		split = strings.Split(s, "_")
	}

	for w, word := range split { // loop on words
		for index, char := range word { // loop on characters

			if index == 0 && w != 0 { // first letter
				if unicode.IsUpper(char) {
					result = result + string(char)
				} else if index == 0 {
					result = result + strings.ToUpper(string(char))
				}
			} else {
				result = result + string(char)
			}
		}
	}

	return result
}
