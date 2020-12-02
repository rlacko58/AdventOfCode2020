package main

import (
	"day1/utils"
	"fmt"
)

// --- Day 1: Report Repair ---
// https://adventofcode.com/2020/day/1

func task1() {
	values := utils.GetIntegers()
	foundValues := utils.FindSumByTwo(values, 2020)
	fmt.Println("values:", foundValues)
	fmt.Println("Multiplied:", foundValues[0] * foundValues[1])
}

func task2() {
	values := utils.GetIntegers()
	foundValues := utils.FindSum(values, 2020, 3)
	fmt.Println("values:", foundValues)
	fmt.Println("Multiplied:", utils.MultiplyIntArray(foundValues))
}

func main() {
	//task1()
	task2()
}