package main

import (
	"day9/utils"
	"fmt"
)

func task() {
	numbers := utils.GetInput()
	invalidNum := utils.FindInvalidNumber(25, &numbers)
	fmt.Println("Invalid Number:", invalidNum)
	i1, i2 := utils.FindContinousSumIndexes(invalidNum, &numbers)
	n1, n2 := utils.FindSmallAndHighNum(numbers[i1:i2+1])
	fmt.Println("Encryption Weakness:", n1+n2)
}

func main() {
	task()
}