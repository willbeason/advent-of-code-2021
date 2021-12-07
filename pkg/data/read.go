package data

import (
	"io/ioutil"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(bytes), "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return lines, nil
}
