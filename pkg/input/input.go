package input

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

func ReadStdinLines() []string {
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

func ReadFileLines(path string) []string {
	values := make([]string, 0)

	byt, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	strArr := strings.Split(string(byt), "\n")
	for _, val := range strArr {
		if (val == "") {
			continue
		}
		values = append(values, val)
	}

	return values
}