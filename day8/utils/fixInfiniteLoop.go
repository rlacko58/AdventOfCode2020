package utils

import (
	"day8/model"
	"log"
)

func createJumpUpdaterFunc(opArr *[]model.Operation, from, to string) func() bool {
	updateIndex := -1
	f := func() bool {
		if updateIndex == -1 {
			updateIndex++
		} else if (*opArr)[updateIndex].Name == from {
			(*opArr)[updateIndex].Name = to
			updateIndex++
		}
		for updateIndex < len(*opArr) && 
			(*opArr)[updateIndex].Name != to {
				updateIndex++
			}
		if (updateIndex >= len(*opArr)) {
			return false
		} 
		(*opArr)[updateIndex].Name = from
		return true
	}
	return f
}

func FixInfiniteLoop(opArr *[]model.Operation) int {
	runAmount := make([]int, len(*opArr))
	accVal := 0
	index := 0
	isJmpUpdate := true
	updateJumpToNop := createJumpUpdaterFunc(opArr, "jmp", "nop")
	updateNopToJump := createJumpUpdaterFunc(opArr, "nop", "jmp")
	for index < len(*opArr) {
		runAmount[index] ++
		if arrValIsMoreThanOne(runAmount) {
			runAmount = make([]int, len(*opArr))
			accVal = 0
			index = 0
			if isJmpUpdate {
				if !updateJumpToNop() {
					isJmpUpdate = false
				}
			} else {
				if !updateNopToJump() {
					log.Fatal("There is no fix")
				}
			}
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