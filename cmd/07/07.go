package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/willbeason/advent-of-code-2021/pkg/data"
)

func main() {
	lines, err := data.ReadLines("cmd/07/data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	crabs, err := readCrabs(lines[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	left := int64(math.MaxInt)
	right := int64(0)

	for p := range crabs {
		if p < left {
			left = p
		}

		if p > right {
			right = p
		}
	}

	minFuel := int64(math.MaxInt64)

	for p := left; p <= right; p++ {
		fuel := crabs.Cost(p)
		if fuel < minFuel {
			minFuel = fuel
		}
	}

	fmt.Println(minFuel)

	minFuel2 := int64(math.MaxInt64)

	for p := left; p <= right; p++ {
		fuel := crabs.Cost2(p)
		if fuel < minFuel2 {
			minFuel2 = fuel
		}
	}

	fmt.Println(minFuel2)
}

type Crabs map[int64]int64

func readCrabs(s string) (Crabs, error) {
	result := Crabs{}

	for _, f := range strings.Split(s, ",") {
		position, err := strconv.ParseInt(f, 10, 64)
		if err != nil {
			return nil, err
		}
		result[position]++
	}

	return result, nil
}

func (c Crabs) Cost(i int64) int64 {
	var cost int64

	for p, n := range c {
		diff := p - i

		if diff < 0 {
			diff = -diff
		}

		cost += n * diff
	}

	return cost
}

func (c Crabs) Cost2(i int64) int64 {
	var cost int64

	for p, n := range c {
		diff := p - i

		if diff < 0 {
			diff = -diff
		}

		cost += n * diff * (diff + 1) / 2
	}

	return cost
}
