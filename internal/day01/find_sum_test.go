package day01

import (
	"aoc-2020/pkg/array"
	"aoc-2020/pkg/convert"
	"aoc-2020/pkg/input"
	"testing"
)

var t1i2Sol = 567171
var t2i2Sol = 212428694

func TestFindSum(t *testing.T) {
	strArr := input.ReadFileLines("../../assets/inputs/day01/day01_2.txt")
	values := convert.StrArrToInt(&strArr)
	solution := FindSum(values, 2020, 2)
	ans := array.ProductOfInts(&solution)
	if ans != t1i2Sol {
		t.Errorf("FindSum(day01_2.txt, 2020, 3) = %d, want %d", ans, t1i2Sol)
	}
	solution = FindSum(values, 2020, 3)
	ans = array.ProductOfInts(&solution)
	if ans != t2i2Sol {
		t.Errorf("FindSum(day01_2.txt, 2020, 3) = %d, want %d", ans, t2i2Sol)
	}
}