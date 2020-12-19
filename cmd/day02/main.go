package main

import (
	"aoc-2020/internal/day02"
	"aoc-2020/pkg/input"
	"fmt"
	"os"
)

// --- Day 2: Password Philosophy ---
// https://adventofcode.com/2020/day/2

func main() {
	strArr := input.ReadLines(os.Stdin)
	policies := day02.ConvertArrToPolicy(&strArr)
	fmt.Println("Task1")
	fmt.Println("Valid passwords:", 
		day02.CountValidPwInRange(day02.ValidatePwBeetweenRange, &policies))
	fmt.Println("\nTask2")
	fmt.Println("Valid passwords:", 
		day02.CountValidPwInRange(day02.ValidatePwLetterOnlyOnce, &policies))
}