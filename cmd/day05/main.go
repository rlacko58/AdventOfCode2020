package main

import (
	"aoc-2020/internal/day05"
	"aoc-2020/pkg/input"
	"fmt"
	"os"
)

// --- Day 5: Binary Boarding ---
// https://adventofcode.com/2020/day/5

func main() {
	strArr := input.ReadLines(os.Stdin)
	fmt.Println("Task1")
	fmt.Println("Highest Seat Id:",
		day05.HighestSeatId(&strArr))

	fmt.Println("\nTask2")
	fmt.Println("Free Seat Id:", 
		day05.FreeSeatId(&strArr))
}