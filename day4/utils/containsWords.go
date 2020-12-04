package utils

import (
	"strings"
)

func ContainsWords(text string, words []string) bool {
	for _, word := range words {
		if ! strings.Contains(text, word) {
			return false
		} 
	}
	return true
}