package main

import (
	"day5/model"
	"day5/utils"
	"fmt"
	"sort"
)

func task1() {
	inputs := utils.GetInput()
	maxSeatId := 0
	for _, str := range(inputs) {
		pos := model.MakeSeatPos(str)
		if seatId := pos.GetSeatId(); seatId > maxSeatId {
			maxSeatId = seatId
		}
	}
	fmt.Println("Highest Seat Id:", maxSeatId)
}

func task2() {
	inputs := utils.GetInput()
	seatIds := make([]int, 0)
	for _, str := range(inputs) {
		pos := model.MakeSeatPos(str)
		seatIds = append(seatIds, pos.GetSeatId())
	}
	sort.Ints(seatIds)
	for i := 1; i < len(seatIds); i++ {
		if seatIds[i] - seatIds[i-1] == 2 {
			fmt.Println("Free Seat Id: ", seatIds[i]-1)
		}
	}
}

func main() {
	// task1()
	task2()
}