package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/willbeason/advent-of-code-2021/pkg/data"
)

func main() {
	lines, err := data.ReadLines("cmd/06/data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fish, err := readFish(lines[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < 80; i++ {
		fish.Age()
	}

	fmt.Println(fish.Population())

	fish, err = readFish(lines[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i < 256; i++ {
		fish.Age()
	}

	fmt.Println(fish.Population())
}

func readFish(s string) (Fish, error) {
	result := make(Fish, 9)

	for _, f := range strings.Split(s, ",") {
		timer, err := strconv.ParseInt(f, 10, 64)
		if err != nil {
			return nil, err
		}
		result[timer]++
	}

	return result, nil
}

type Fish []int

func (f *Fish) Age() {
	babies := (*f)[0]
	*f = (*f)[1:]
	(*f)[6] += babies
	*f = append(*f, babies)
}

func (f *Fish) Population() int {
	pop := 0
	for _, p := range *f {
		pop += p
	}

	return pop
}
