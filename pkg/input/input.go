package input

import (
	"bufio"
	"io"
)

func ReadLines(r io.Reader) []string {
	values := make([]string, 0)

	scanner := bufio.NewScanner(r)

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