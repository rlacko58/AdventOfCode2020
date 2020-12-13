package utils

import (
	"log"
	"regexp"
	"strconv"
)

func ConvertBusString(busStr string) []int {
	re := regexp.MustCompile(`((?:\d+)+),*`)
	match := re.FindAllStringSubmatch(busStr, -1)
	busIds := make([]int, 0)
	for _, v := range(match) {
		num, err := strconv.Atoi(v[1])
		if err != nil {
			log.Fatal(err)
		} 
		busIds = append(busIds, num)
	}
	return busIds
} 