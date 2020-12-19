package main

import (
	"aoc-2020/internal/day04"
	"fmt"
	"os"
)

// --- Day 4: Passport Processing ---
// https://adventofcode.com/2020/day/4

func main() {
	strArr := day04.GetPassports(os.Stdin)
	fmt.Println("Task1")
	fmt.Println("Valid Passports:",
		day04.CountValidPass(day04.ContainsWords, &strArr))

	fmt.Println("\nTask2")
	fmt.Println("Valid Passports:", 
		day04.CountValidPass(day04.IsValidPass, &strArr))
}