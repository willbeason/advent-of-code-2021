package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/willbeason/advent-of-code-2021/pkg/data"
)

func main() {
	lines, err := data.ReadLines("cmd/02/data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rows := make([]row, len(lines))
	for i := range rows {
		rows[i], err = read(lines[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	var (
		depth      int64
		horizontal int64
		aim        int64
	)

	for _, r := range rows {
		switch r.direction {
		case forward:
			horizontal += r.distance
		case down:
			depth += r.distance
		case up:
			depth -= r.distance
		}
	}

	fmt.Println(depth * horizontal)

	depth = 0
	horizontal = 0
	aim = 0

	for _, r := range rows {
		switch r.direction {
		case forward:
			horizontal += r.distance
			depth += r.distance * aim
		case down:
			aim += r.distance
		case up:
			aim -= r.distance
		}
	}

	fmt.Println(depth * horizontal)
}

type Direction string

const (
	forward = "forward"
	down    = "down"
	up      = "up"
)

type row struct {
	direction Direction
	distance  int64
}

func read(line string) (row, error) {
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		return row{}, fmt.Errorf("%w: got %d parts in %q",
			data.ErrInvalidData, len(parts), line)
	}

	r := row{}

	switch parts[0] {
	case forward, down, up:
		r.direction = Direction(parts[0])
	default:
		return row{}, fmt.Errorf("%w: unrecognized direction %q in %q",
			data.ErrInvalidData, parts[0], line)
	}

	var err error
	r.distance, err = strconv.ParseInt(parts[1], 10, 64)

	return r, err
}
