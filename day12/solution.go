package main

import (
	"day12/utils"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
)

type Location struct {
	Facing, X, Y int
}

type Instruction struct {
	action string
	value int
}

func (loc *Location) rotate(dir string, degrees int) {
	// degrees % 90 == 0
	if dir == "R" { 
		degrees = 4*90 - degrees 
	}
	loc.Facing = (loc.Facing + degrees / 90) % 4
}

func convertDirIntToStr(dir int) string {
	switch (dir) {
	case 0: 
		return "E"
	case 1:
		return "N"
	case 2:
		return "W"
	default:
		return "S"
	}
}

func (loc *Location) move (dir string, value int) {
	switch(dir) {
	case "N":
		loc.Y += value
	case "S":
		loc.Y -= value
	case "E":
		loc.X += value
	case "W":
		loc.X -= value
	case "L":
		loc.rotate("L", value)
	case "R":
		loc.rotate("R", value)
	case "F":
		loc.move(convertDirIntToStr(loc.Facing), value)
	}
}

func convertToInstruction(instr *string) Instruction {
	re := regexp.MustCompile(`^(\w{1})(\d+)`)
	match := re.FindStringSubmatch(*instr)
	val, err := strconv.Atoi(match[2]) 
	if err != nil {
		log.Fatal(err)
	}
	return Instruction{action: match[1], value: val}
}

func abs(val int) int {
	return int(math.Abs(float64(val)))
}

func task1(instr *[]string) {
	instrArr := make([]Instruction, 0)
	loc := Location{}
	for _, v := range(*instr) {
		instrArr = append(instrArr, convertToInstruction(&v))
	}
	for _, v := range(instrArr) {
		loc.move(v.action, v.value)
	}
	fmt.Println(abs(loc.X) + abs(loc.Y))
}

func main() {
	instr := utils.GetInput()
	task1(&instr)
}