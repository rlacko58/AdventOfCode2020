package utils

import (
	"sort"
)

func FindSmallAndHighNum(numbers []int) (int, int) {
	sort.Ints(numbers)
	return numbers[0], numbers[len(numbers)-1]
}