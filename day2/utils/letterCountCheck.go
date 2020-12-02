package utils

import "day2/model"

func LetterCountCheck(value model.PolicyCheck) bool {
	letterCounts := CountCharacters(value.Password)
	for key, element := range letterCounts {
		if key == value.Letter {
			return element >= value.From && element <= value.To
		}
	}
	return value.From == 0
}