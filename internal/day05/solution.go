package day05

import (
	"sort"
)

func HighestSeatId(strArr *[]string) int {
	maxSeatId := 0
	for _, str := range(*strArr) {
		pos := MakeSeatPos(str)
		if seatId := pos.GetSeatId(); seatId > maxSeatId {
			maxSeatId = seatId
		}
	}
	return maxSeatId
}

func FreeSeatId(strArr *[]string) int {
	seatIds := make([]int, 0)
	for _, str := range(*strArr) {
		pos := MakeSeatPos(str)
		seatIds = append(seatIds, pos.GetSeatId())
	}
	sort.Ints(seatIds)
	for i := 1; i < len(seatIds); i++ {
		if seatIds[i] - seatIds[i-1] == 2 {
			return seatIds[i] - 1
		}
	}
	return -1
}