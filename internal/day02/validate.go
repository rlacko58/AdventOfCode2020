package day02

type validatePw func(*Policy) bool

func countCharacters(text *string) map[string]int {
	m := make(map[string]int)
	for _, char := range(*text) {
		m[string(char)] += 1
	}
	return m
}

func ValidatePwBeetweenRange(pwd *Policy) bool {
	letterCounts := countCharacters(&pwd.Password)
	for key, count := range letterCounts {
		if key == pwd.Letter {
			return count >= pwd.From && count <= pwd.To
		}
	}
	return pwd.From == 0
}

func ValidatePwLetterOnlyOnce(pwd *Policy) bool {
	firstL, secondL := string(pwd.Password[pwd.From-1]), 
					string(pwd.Password[pwd.To-1])
	return (firstL == pwd.Letter || secondL == pwd.Letter ) && 
		(firstL != secondL)
}

func CountValidPwInRange(pwCheckFunc validatePw, pwds *[]Policy) int {
	validPwNum := 0
	for _, val := range(*pwds) {
		if pwCheckFunc(&val) {
			validPwNum++
		}
	}
	return validPwNum
}
