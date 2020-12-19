package main

import (
	"aoc-2020/internal/day05"
	"aoc-2020/pkg/input"
	"os"
	"testing"
)


func TestDay05(t *testing.T) {
	var t1i2Sol = 828
	var t2i2Sol = 565

	file, err := os.Open("data/day05_2.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	strArr := input.ReadLines(file)
	
	ans := day05.HighestSeatId(&strArr)
	if ans != t1i2Sol {
		t.Errorf("HighestSeatId() = %d, want %d", ans, t1i2Sol)
	}
	ans = day05.FreeSeatId(&strArr)
	if ans != t2i2Sol {
		t.Errorf("FreeSeatId() = %d, want %d", ans, t2i2Sol)
	}
}