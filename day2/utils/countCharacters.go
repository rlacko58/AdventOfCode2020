package utils

func CountCharacters(text string) map[string]int {
	m := make(map[string]int)
	for _, char := range text {
		m[string(char)] += 1
	}
	return m
}