package main

import (
	"day13/utils"
	"fmt"
	"log"
	"strconv"
)

func task1(inp *[]string) {
	busIds := utils.ConvertBusString((*inp)[1])
	currTime, err := strconv.Atoi((*inp)[0])
	if err != nil {
		log.Fatal(err)
	}
	nextDepart, busId := utils.FindClosestNum(currTime, &busIds)
	fmt.Println("Task1:", (nextDepart-currTime)*busId)
}

func task2(inp *[]string) {
	busIds := utils.ConvertNewBusString((*inp)[1])
	fmt.Println("Task2:",utils.FindNumber(&busIds))
}

func main() {
	inp := utils.GetInput()
	task1(&inp)
	task2(&inp)
}