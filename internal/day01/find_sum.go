package day01

import "aoc-2020/pkg/array"

func arrayValuesLessThan(array []int, value int) bool {
	for _, elem := range array {
		if elem >= value {
			return false
		}
	}
	return true
}

func incrementArrayValues(array []int, max, index int) {
	if index >= len(array) {
		return
	}
	array[index] += 1
	if array[index] >= max {
		array[index] = 0
		incrementArrayValues(array, max, index + 1)
	}
}

func FindSum(values []int, target, pairNum int) []int {
	i := make([]int, pairNum)
	for arrayValuesLessThan(i, len(values)) {
		finalArr := make([]int, 0)
		for j := 0; j < pairNum; j++ {
			finalArr = append(finalArr, values[i[j]])
		}
		if array.SumOfInts(&finalArr) == target {
			return finalArr
		}
		incrementArrayValues(i, len(values), 0)
	}
	return nil
}