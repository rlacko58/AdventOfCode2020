package main

import (
	"day3/utils"
	"fmt"
)

// --- Day 3: Toboggan Trajectory ---
// https://adventofcode.com/2020/day/3

func task1() {
	rows := utils.GetInput()
	fmt.Println("number of encountered trees:", 
		utils.RightThreeDownOne(rows))
}

func task2() {
	rows := utils.GetInput()
	values := make([]int, 0)
	values = append(values, utils.Sloope(rows, 1, 1))
	values = append(values, utils.Sloope(rows, 3, 1))
	values = append(values, utils.Sloope(rows, 5, 1))
	values = append(values, utils.Sloope(rows, 7, 1))
	values = append(values, utils.Sloope(rows, 1, 2))

	fmt.Println("number of encountered trees:", 
		values[0],
		values[1],
		values[2],
		values[3],
		values[4],
	)
	multiplied := 1
	for _, value := range(values) {
		multiplied *= value
	}
	fmt.Println("multiplied:", multiplied)
	
}

func main() {
	// task1()
	task2()
}