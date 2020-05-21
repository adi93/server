package utils

import (
	"encoding/json"
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

// PrintStruct prints a structure in a pretty manner.
func PrintStruct(s interface{}) string {
	str, err := json.Marshal(s)
	if err != nil {
		panic("Could not covert into json format")
	}
	return string(str)
}
