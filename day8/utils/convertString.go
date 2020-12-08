package utils

import (
	"log"
	"regexp"
	"strconv"
)

func ConvertString(str string) (string, int) {
	re := regexp.MustCompile(`(\w+)\s(\+|\-)(\d+)`)
	match := re.FindStringSubmatch(str)
	number, err := strconv.Atoi(match[3])
	if err != nil {
		log.Fatal(err)
	}
	if match[2] == "-" {
		return match[1], -number
	}
	return match[1], number
}