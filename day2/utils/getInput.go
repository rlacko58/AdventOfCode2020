package utils

import (
	"bufio"
	"day2/model"
	"log"
	"os"
	"strconv"
	"strings"
)

func getRange(text string) (int, int) {
	numbers := strings.Split(text, "-")
	rangeNumbers := make([]int, 0)
	for _, val := range numbers {
		number, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		rangeNumbers = append(rangeNumbers, number)
	}
	return rangeNumbers[0], rangeNumbers[1]
}

func GetInput() []model.PolicyCheck {
	values := make([]model.PolicyCheck, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			parts := strings.Split(text, " ")
			from, to := getRange(parts[0])
			letter := strings.Replace(parts[1], ":", "", 1)
			
			values = append(values, model.PolicyCheck{
				From: from,
				To: to,
				Letter: letter,
				Password: parts[2],
			})
		} else {
			break
		}
	}

	return values
}