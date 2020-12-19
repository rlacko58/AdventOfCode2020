package main

import (
	"aoc-2020/internal/day01"
	"aoc-2020/pkg/array"
	"aoc-2020/pkg/convert"
	"aoc-2020/pkg/input"
	"os"
	"testing"
)

func TestDay01(t *testing.T) {
	var t1i2Sol = 567171
	var t2i2Sol = 212428694

	file, err := os.Open("data/day01_2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	strArr := input.ReadLines(file)
	
	values := convert.StrArrToInt(&strArr)
	solution := day01.FindSum(values, 2020, 2)
	ans := array.ProductOfInts(&solution)
	if ans != t1i2Sol {
		t.Errorf("ProductOfInts(day01_2.txt, 2020, 3) = %d, want %d", ans, t1i2Sol)
	}
	solution = day01.FindSum(values, 2020, 3)
	ans = array.ProductOfInts(&solution)
	if ans != t2i2Sol {
		t.Errorf("ProductOfInts(day01_2.txt, 2020, 3) = %d, want %d", ans, t2i2Sol)
	}
}