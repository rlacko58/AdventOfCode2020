package main

import (
	"aoc-2020/internal/day02"
	"aoc-2020/pkg/input"
	"testing"
)


func TestDay02(t *testing.T) {
	var t1i2Sol = 447
	var t2i2Sol = 249

	strArr := input.ReadFileLines("data/day02_2.txt")
	policies := day02.ConvertArrToPolicy(&strArr)
	ans := day02.CountValidPwInRange(day02.ValidatePwBeetweenRange, &policies)
	if ans != t1i2Sol {
		t.Errorf("ValidatePwBeetweenRange(day02_2.txt) = %d, want %d", ans, t1i2Sol)
	}
	ans = day02.CountValidPwInRange(day02.ValidatePwLetterOnlyOnce, &policies)
	if ans != t2i2Sol {
		t.Errorf("ValidatePwLetterOnlyOnce(day02_2.txt) = %d, want %d", ans, t2i2Sol)
	}
}