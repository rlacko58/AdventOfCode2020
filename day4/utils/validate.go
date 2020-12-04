package utils

import (
	"strconv"
)

func isFourDigitNumber(text string) bool {
	if len(text) != 4 {
		return false
	}
	if _, err := strconv.Atoi(text); err != nil {
		return false
	}
	return true
}
