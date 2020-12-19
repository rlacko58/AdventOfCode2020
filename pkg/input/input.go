package input

import (
	"bufio"
	"os"
)

func GetStrings() []string {
	values := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			values = append(values, text)
		} else {
			break
		}
	}

	return values
}