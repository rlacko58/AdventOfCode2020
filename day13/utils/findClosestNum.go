package utils

func FindClosestNum(curr int, ids *[]int) (int, int) {
	num := curr
	for {
		for _, v :=  range(*ids) {
			if num % v == 0 {
				return num, v
			}
		}
		num++
	}
}