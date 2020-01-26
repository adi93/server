package utils

import (
	"unicode"
)

// IsBlank checks if a string is empty or is made of white space characters
func IsBlank(s string) bool {
	if s == "" {
		return true
	}

	for _, char := range s {
		if unicode.IsSpace(char) {
			return false
		}
	}
	return true
}
