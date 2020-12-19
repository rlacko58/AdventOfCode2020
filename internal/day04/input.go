package day04

import (
	"bufio"
	"io"
)

func GetPassports(r io.Reader) []string {
	values := make([]string, 0)

	scanner := bufio.NewScanner(r)

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