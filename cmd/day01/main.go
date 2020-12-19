package main

import (
	"aoc-2020/internal/day01"
	"aoc-2020/pkg/convert"
	"aoc-2020/pkg/input"
	"fmt"
)

// --- Day 1: Report Repair ---
// https://adventofcode.com/2020/day/1

func main() {
	strArr := input.GetStrings()
	values := convert.StrArrToInt(&strArr)
	task1Nums := day01.FindSum(values, 2020, 2)
	task2Nums := day01.FindSum(values, 2020, 3)
	fmt.Println("## Task1")
	day01.Output(&task1Nums)
	fmt.Println("\n## Task 2")
	day01.Output(&task2Nums)
}