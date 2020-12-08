package utils

import (
	"day8/model"
)

func ConvertStringArr(strArr *[]string) []model.Operation {
	opArr := make([]model.Operation, 0)
	for _, str := range(*strArr) {
		op, num := ConvertString(str)
		opArr = append(opArr, model.Operation{Name: op, Num: num})
	}
	return opArr
}