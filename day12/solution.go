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

type Waypoint struct {
	X, Y int
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

func (loc *Waypoint) rotate(dir string, degrees int) {
	// degrees % 90 == 0
	if dir == "R" { 
		degrees = 4*90 - degrees 
	}
	rotateLeftNum := (degrees / 90) % 4
	for i := 0; i < rotateLeftNum; i++ {
		tmp := loc.X
		loc.X = - loc.Y
		loc.Y = tmp
	}
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

func (loc *Waypoint) move (dir string, value int) {
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

func convertToInstr(instr *[]string) []Instruction {
	instrArr := make([]Instruction, 0)
	for _, v := range(*instr) {
		instrArr = append(instrArr, convertToInstruction(&v))
	}
	return instrArr
}

func task1(instr *[]string) {
	instrArr := convertToInstr(instr)
	loc := Location{}
	for _, v := range(instrArr) {
		loc.move(v.action, v.value)
	}
	fmt.Println("Task1 Location:", abs(loc.X) + abs(loc.Y))
}

func task2(instr *[]string) {
	instrArr := convertToInstr(instr)
	shipLoc := Location{}
	wpLoc := Waypoint{X: 10, Y: 1}
	for _, v := range(instrArr) {
		if (v.action == "F") {
			shipLoc.X += v.value * wpLoc.X
			shipLoc.Y += v.value * wpLoc.Y
		} else {
			wpLoc.move(v.action, v.value)
		}
	}
	fmt.Println("Task2 Location:", abs(shipLoc.X) + abs(shipLoc.Y))
}

func main() {
	instr := utils.GetInput()
	task1(&instr)
	task2(&instr)
}