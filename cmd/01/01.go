package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("cmd/01/data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	increased := 0

	depths, err := toInts(bytes)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	previous := depths[0]

	for i := 1; i < len(depths); i++ {
		next := depths[i]
		if next > previous {
			increased++
		}

		previous = next
	}

	fmt.Println(increased)

	increasedSliding := 0

	n0, n1, n2 := depths[0], depths[1], depths[2]
	for i := 3; i < len(depths); i++ {
		n3 := depths[i]
		if n0 < n3 {
			increasedSliding++
		}
		n0, n1, n2 = n1, n2, n3
	}

	fmt.Println(increasedSliding)
}

func toInts(bytes []byte) ([]int64, error) {
	lines := strings.Split(string(bytes), "\n")

	var err error
	result := make([]int64, len(lines))
	for i, line := range lines {
		result[i], err = readLine(line)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func readLine(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
