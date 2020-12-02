package utils

func FindSumByTwo(values []int, target int) []int {
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values); j++ {
			if target - values[i] == values[j] {
				return []int{values[i], values[j]}
			}
		}
	}
	return nil
}