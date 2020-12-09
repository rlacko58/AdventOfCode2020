package utils

func FindSumInRange(sum, from, to int, numbers *[]int) bool {
	for i := from; i < to; i++ {
		for j := from; j < to; j++ {
			if i == j {
				continue
			}
			if (*numbers)[i] + (*numbers)[j] == sum {
				return true
			}
		}
	}
	return false
}