package utils

func MultiplyIntArray(array []int) int {
	result := 1
	for _, val := range array {
		result *= val
	}
	return result
}