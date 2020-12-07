package main

import (
	"day7/model"
	"day7/utils"
	"fmt"
)

func part1() {
	strs := utils.GetInput()
	bags := model.CreateBags(strs)
	fmt.Println("Container bags:",
		utils.CountParentContainers("shiny gold", bags))
}

func part2() {
	strs := utils.GetInput()
	bags := model.CreateBags(strs)
	fmt.Println("Container bags:",
		utils.NumberOfBags("shiny gold", bags))
}

func main() {
	// part1()
	part2()
}