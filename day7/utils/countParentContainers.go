package utils

import (
	"day7/model"
)

func countParrent(currBag *model.Bag, containers map[string]bool) {
	containers[currBag.Color] = true

	for _, v := range(currBag.ParentBags) {
		countParrent(v, containers)
	}
}

func CountParentContainers(color string, bags []*model.Bag) int {
	bag := model.GetBag(&color, bags)
	containerBags := make(map[string]bool)
	countParrent(bag, containerBags)
	number := 0
	for _, v := range(containerBags) {
		if v {
			number ++
		}
 	}
	return number -1
}