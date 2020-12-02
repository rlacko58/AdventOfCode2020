package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func GetIntegers() []int {
	values := make([]int, 0)
	
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			number, err := strconv.Atoi(text)
			if err != nil {
				log.Fatal(err)
			}
			values = append(values, number)
		} else {
			break
		}
	}
	return values
}