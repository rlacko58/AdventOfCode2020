package utils

func CountEveryoneYes(answers string) int {
	uniqueAnswers := make(map[string]int)
	numOfPeople := 1
	for i := 0; i < len(answers); i++ {
		if(string(answers[i]) == " ") {
			numOfPeople++
			continue
		}
		uniqueAnswers[string(answers[i])]++
	}
	number := 0
	for _,answ := range(uniqueAnswers) {
		if(answ == numOfPeople) {
			number++
		}
	}
	return number
}