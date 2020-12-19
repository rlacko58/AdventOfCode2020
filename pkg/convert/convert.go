package convert

import (
	"log"
	"strconv"
)

func ToInt(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return number
}