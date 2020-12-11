package main

import (
	"day10/utils"
	"fmt"
	"sort"
)

func task1() {
	numbers := utils.GetInput()
	numbers = append([]int{0}, numbers...)
	diffs := make([]int, 0)
	sort.Ints(numbers)
	for i := 0; i < len(numbers)-1; i++  {
		diffs = append(diffs, numbers[i+1]-numbers[i])
	}
	diffs = append(diffs, 3)
	countDiffs := make(map[int]int)
	for _, v := range(diffs) {
		countDiffs[v] ++
	}
	fmt.Println("First task solution:", countDiffs[1]*countDiffs[3])
}

func walk(arr *[]int, index int) int {
	if index == len(*arr) - 1 {
		return 1
	}
	num := 0

	for i := 1; i <= 3; i++ {
		if index + i == len(*arr) {
			break
		}
		if (*arr)[index+i] - (*arr)[index] <= 3 {
			num += walk(arr, index+i)
		}
	}
	return num
}

func task2() {
	numbers := utils.GetInput()
	numbers = append([]int{0}, numbers...)
	sort.Ints(numbers)
	numbers = append(numbers, numbers[len(numbers)-1]+3)
	diffs := make([]int, 0)
	for i := range(numbers) {
		if i == len(numbers) - 1 {
			break
		}
		diffs = append(diffs, numbers[i+1]-numbers[i])
	}
	oneCounts := make([]int, 0)
	count := 0
	for i := 0; i < len(diffs); i++ {
		if diffs[i] == 1 {
			count++
		} else if diffs[i] == 3 {
			if count == 0 {
				continue
			}
			oneCounts = append(oneCounts, count)
			count = 0
		} else {
			fmt.Println("It should not contain 2")
		}
	}
	sum := 1
	for _, v := range(oneCounts) {
		if v == 2 {
			sum *= 2
		} else if v > 2 {
			sum *= (4 + 3 * (v-3) )
		}
	}
	fmt.Println("Second task solution:",sum)
}

func main() {
	// task1()
	task2()
}