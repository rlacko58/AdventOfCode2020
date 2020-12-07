package utils

import (
	"day7/model"
)

func getBagContainsNum(bag *model.Bag, num *int) {
	*num ++
	for _, v := range(bag.Contains) {
		for i := 0; i< v.Number; i++ {
			getBagContainsNum(v.Bag, num)
		}
	}
}

func NumberOfBags(color string, bags []*model.Bag) int {
	bag := model.GetBag(&color, bags)

	num := 0
	getBagContainsNum(bag, &num)

	return num - 1
}