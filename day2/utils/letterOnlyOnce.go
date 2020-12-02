package utils

import "day2/model"

func LetterOnlyOnce(value model.PolicyCheck) bool {
	letter := value.Letter
	firstLetter := string(value.Password[value.From-1])
	secondLetter := string(value.Password[value.To-1])
	return (firstLetter == letter || secondLetter == letter) && 
		(firstLetter != secondLetter)
}
