package main

import (
	"day8/utils"
	"fmt"
)

func task1() {
	strArr := utils.GetInput()
	opArr := utils.ConvertStringArr(&strArr)
	fmt.Println("Acc value:", utils.CalcAcc(&opArr))
}

func task2() {
	strArr := utils.GetInput()
	opArr := utils.ConvertStringArr(&strArr)
	fmt.Println("Acc value:", utils.FixInfiniteLoop(&opArr))
}

func main() {
	// task1()
	task2()
}