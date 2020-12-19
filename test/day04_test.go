package main

import (
	"aoc-2020/internal/day04"
	"os"
	"testing"
)


func TestDay04(t *testing.T) {
	var t1i2Sol = 222
	var t2i2Sol = 140

	file, err := os.Open("data/day04_2.txt")
	if err != nil {
		panic(err)
	}
	rows := day04.GetPassports(file)
	ans := day04.CountValidPass(day04.ContainsWords, &rows)
	if ans != t1i2Sol {
		t.Errorf("ContainsWords() = %d, want %d", ans, t1i2Sol)
	}
	ans = day04.CountValidPass(day04.IsValidPass, &rows)
	if ans != t2i2Sol {
		t.Errorf("IsValidPass() = %d, want %d", ans, t2i2Sol)
	}
}