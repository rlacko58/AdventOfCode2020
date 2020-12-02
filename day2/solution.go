package main

import (
	"day2/utils"
	"fmt"
)

// --- Day 2: Password Philosophy ---
// https://adventofcode.com/2020/day/2

func task1() {
	values := utils.GetInput()
	validPasswordNum := 0
	for _, value := range values {
		if utils.LetterCountCheck(value) {
			validPasswordNum++
		}
	}
	fmt.Println("Valid passwords:", validPasswordNum)
}

func task2() {
	values := utils.GetInput()
	validPasswordNum := 0
	for _, value := range values {
		if utils.LetterOnlyOnce(value) {
			validPasswordNum++
		}
	}
	fmt.Println("Valid passwords:", validPasswordNum)
}

func main() {
	// task1()
	task2()
}