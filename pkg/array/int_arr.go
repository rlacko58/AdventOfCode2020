package array

func SumOfInts(array *[]int) int {
	result := 0
	for _, val := range(*array) {
		result += val
	}
	return result
}

func ProductOfInts(array *[]int) int {
	result := 1
	for _, val := range(*array) {
		result *= val
	}
	return result
}
	