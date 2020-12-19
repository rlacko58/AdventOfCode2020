package day01

import (
	"aoc-2020/pkg/array"
	"fmt"
)


func Output(values *[]int) {
	fmt.Println("values:", *values)
	fmt.Println("Product:", array.ProductOfInts(values))
}