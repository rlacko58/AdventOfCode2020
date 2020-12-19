package main

import (
	"aoc-2020/internal/day03"
	"aoc-2020/pkg/array"
	"aoc-2020/pkg/input"
	"testing"
)


func TestDay03(t *testing.T) {
	var t1i2Sol = 250
	var t2i2Sol = 1592662500

	rows := input.ReadFileLines("data/day03_2.txt")
	ans := day03.CountEncounteredTrees(&rows, 3, 1)
	if ans != t1i2Sol {
		t.Errorf("CountEncounteredTrees(day03_2.txt, 3, 1) = %d, want %d", ans, t1i2Sol)
	}
	values := make([]int, 0)
	values = append(values, day03.CountEncounteredTrees(&rows, 1, 1))
	values = append(values, day03.CountEncounteredTrees(&rows, 3, 1))
	values = append(values, day03.CountEncounteredTrees(&rows, 5, 1))
	values = append(values, day03.CountEncounteredTrees(&rows, 7, 1))
	values = append(values, day03.CountEncounteredTrees(&rows, 1, 2))
	ans = array.ProductOfInts(&values)
	if ans != t2i2Sol {
		t.Errorf("ProductOfInts(rows) = %d, want %d", ans, t2i2Sol)
	}
}