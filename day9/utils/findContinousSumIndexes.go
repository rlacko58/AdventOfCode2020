package utils

func FindContinousSumIndexes(goalSum int, numbers *[]int)(int, int) {
	for i := 0; i < len(*numbers); i++ {
		sum := 0
		j := i
		for j < len(*numbers) && sum < goalSum {
			sum += (*numbers)[j]
			j++
		}   
		if sum == goalSum {
			return i, j-1
		}
	}
	return -1, -1
}