package main

import (
	"fmt"
	"os"

	"github.com/willbeason/advent-of-code-2021/pkg/data"
)

func main() {
	lines, err := data.ReadLines("cmd/11/data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	octopuses := readOctopuses(lines)

	nFlashes := 0
	for i := 0; i < 100; i++ {
		nFlashes += octopuses.Step()
	}

	fmt.Println(nFlashes)

	for i := 0; i < 1000; i++ {
		sync := octopuses.Step()
		if sync == 100 {
			fmt.Println(i + 101)
			break
		}
	}
}

func readOctopuses(lines []string) *Octopuses {
	o := &Octopuses{
		Energy: make([]int8, 100),
	}

	for i, line := range lines {
		for j, c := range line {
			o.Energy[i*10+j] = int8(c - '0')
		}
	}

	return o
}

type Octopuses struct {
	Energy []int8
}

func (o *Octopuses) Step() int {
	count := 0

	for i := range o.Energy {
		o.Energy[i]++
	}

	flashed := make([]bool, len(o.Energy))

	for {
		nFlashed := 0

		for i, e := range o.Energy {
			if flashed[i] {
				continue
			}

			if e > 9 {
				o.flash(i)

				flashed[i] = true

				nFlashed++
			}
		}

		if nFlashed == 0 {
			break
		}

		count += nFlashed
	}

	for i, f := range flashed {
		if f {
			o.Energy[i] = 0
		}
	}

	return count
}

func (o *Octopuses) flash(i int) {
	notLeft := i%10 > 0
	notRight := i%10 < 9
	notTop := i/10 > 0
	notBottom := i/10 < 9

	if notTop && notLeft {
		o.Energy[i-11]++
	}

	if notTop {
		o.Energy[i-10]++
	}

	if notTop && notRight {
		o.Energy[i-9]++
	}

	if notLeft {
		o.Energy[i-1]++
	}

	if notRight {
		o.Energy[i+1]++
	}

	if notBottom && notLeft {
		o.Energy[i+9]++
	}

	if notBottom {
		o.Energy[i+10]++
	}

	if notBottom && notRight {
		o.Energy[i+11]++
	}
}
