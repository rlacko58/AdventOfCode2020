package convert

import (
	"strconv"
)

func ToInt(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return number
}