package utils

import (
	"bufio"
	"os"
)

func GetStrings() []string {
	values := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)

	text := ""
	for {
		scanner.Scan()
		scannedText := scanner.Text()
		if len(scannedText) == 0 && len(text) != 0 {
			values = append(values, text[:len(text)-1])
			text = ""
			continue
		} else {
			text += string(scannedText + " ")
		}

		if len(text) == 1 {
			break
		}
	}

	return values
}