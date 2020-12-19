package main

import (
	"aoc-2020/internal/day03"
	"aoc-2020/pkg/array"
	"aoc-2020/pkg/input"
	"fmt"
	"os"
)

// --- Day 3: Toboggan Trajectory ---
// https://adventofcode.com/2020/day/3

func main() {
	rows := input.ReadLines(os.Stdin)
	fmt.Println("Task1")
	fmt.Println("Number of encountered trees:",
		day03.CountEncounteredTrees(&rows, 3, 1))

	fmt.Println("\nTask2")
	values := make([]int, 0)
	values = append(values, day03.CountEncounteredTrees(&rows, 1, 1))
	values = append(values, day03.CountEncounteredTrees(&rows, 3, 1))
	values = append(values, day03.CountEncounteredTrees(&rows, 5, 1))
	values = append(values, day03.CountEncounteredTrees(&rows, 7, 1))
	values = append(values, day03.CountEncounteredTrees(&rows, 1, 2))
	fmt.Println("Number of encountered trees:", values)
	fmt.Println("Product:", array.ProductOfInts(&values))
}