package main

import (
	"day6/utils"
	"fmt"
)

func task1() {
	groups := utils.GetAnswers()
	sum := 0
	for _, groupAnswers := range(groups) {
		sum += utils.CountUniqueAnswers(groupAnswers)
	}
	fmt.Println("Number of Answers:", sum)
}

func task2() {
	groups := utils.GetAnswers()
	sum := 0
	for _, groupAnswers := range(groups) {
		sum += utils.CountEveryoneYes(groupAnswers)
	}
	fmt.Println("Number of Answers:", sum)
}

func main() {
	// task1()
	task2()
}