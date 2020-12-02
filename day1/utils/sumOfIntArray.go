package utils

func SumOfIntArray(array []int) int {
	result := 0
	for _, val := range array {
		result += val
	}
	return result
}