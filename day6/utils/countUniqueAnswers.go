package utils

func CountUniqueAnswers(answers string) int {
	uniqueAnswers := make(map[string]bool)
	for i := 0; i < len(answers); i++ {
		if(string(answers[i]) == " ") {
			continue
		}
		uniqueAnswers[string(answers[i])] = true
	}
	number := 0
	for range(uniqueAnswers) {
		number++;
	}
	return number
}