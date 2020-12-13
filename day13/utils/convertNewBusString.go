package utils

import (
	"log"
	"regexp"
	"strconv"
)

func ConvertNewBusString(busStr string) [][2]int {
	re := regexp.MustCompile(`((?:\w+)+),*`)
	match := re.FindAllStringSubmatch(busStr, -1)
	busIds := make([][2]int, 0)
	time := 0
	for _, v := range(match) {
		if v[1] == "x" {
			time++
			continue
		}
		num, err := strconv.Atoi(v[1])
		if err != nil {
			log.Fatal(err)
		} 
		busIds = append(busIds, [2]int{num, time})
		time++
	}
	return busIds
} 