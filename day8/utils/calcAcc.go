package utils

import (
	"day8/model"
)

func arrValIsMoreThanOne(arr []int) bool {
	for _, v := range(arr) {
		if v > 1 {
			return true
		}
	}
	return false
}

func CalcAcc(opArr *[]model.Operation) int {
	runAmount := make([]int, len(*opArr))
	accVal := 0
	index := 0
	for index < len(*opArr) {
		runAmount[index] ++
		if arrValIsMoreThanOne(runAmount) {
			break
		}
		switch((*opArr)[index].Name) {
		case "nop":
			index++
		case "acc":
			accVal += (*opArr)[index].Num
			index++
		case "jmp":
			index += (*opArr)[index].Num
		}
	}
	return accVal
}