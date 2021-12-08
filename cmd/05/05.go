package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/willbeason/advent-of-code-2021/pkg/data"
)

func main() {
	lines, err := data.ReadLines("cmd/05/data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	vents := make([]Vents, len(lines))

	for i, line := range lines {
		vents[i], err = readLine(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	var (
		xDim int64
		yDim int64
	)

	for _, v := range vents {
		if v.StartX >= xDim {
			xDim = v.StartX + 1
		}

		if v.EndX >= xDim {
			xDim = v.EndX + 1
		}

		if v.StartY >= yDim {
			yDim = v.StartY + 1
		}

		if v.EndY >= yDim {
			yDim = v.EndY + 1
		}
	}

	m := Map{
		XDim:  xDim,
		Vents: make([]int64, xDim*yDim),
	}

	for _, v := range vents {
		if v.IsStraight() {
			m.AddVents(v)
		}
	}

	fmt.Println(m.Score())

	m2 := Map{
		XDim:  xDim,
		Vents: make([]int64, xDim*yDim),
	}

	for _, v := range vents {
		m2.AddVents(v)
	}

	fmt.Println(m2.Score())
}

type Map struct {
	XDim  int64
	Vents []int64
}

func (m *Map) AddVents(v Vents) {
	var (
		dx int64
		dy int64
	)

	if v.StartX < v.EndX {
		dx = 1
	} else if v.StartX > v.EndX {
		dx = -1
	}

	if v.StartY < v.EndY {
		dy = 1
	} else if v.StartY > v.EndY {
		dy = -1
	}

	for x, y := v.StartX-dx, v.StartY-dy; x != v.EndX || y != v.EndY; {
		x += dx
		y += dy

		idx := x + y*m.XDim
		m.Vents[idx]++
	}
}

func (m *Map) Score() int {
	score := 0

	for _, v := range m.Vents {
		if v > 1 {
			score++
		}
	}

	return score
}

type Vents struct {
	StartX, StartY int64
	EndX, EndY     int64
}

func (v Vents) IsStraight() bool {
	return v.StartX == v.EndX || v.StartY == v.EndY
}

func readLine(s string) (Vents, error) {
	parts := strings.Split(s, " -> ")
	if len(parts) != 2 {
		return Vents{}, fmt.Errorf("%w: invalid line: %q",
			data.ErrInvalidData, s)
	}

	start := strings.Split(parts[0], ",")
	if len(start) != 2 {
		return Vents{}, fmt.Errorf("%w: invalid line: %q",
			data.ErrInvalidData, s)
	}

	end := strings.Split(parts[1], ",")
	if len(end) != 2 {
		return Vents{}, fmt.Errorf("%w: invalid line: %q",
			data.ErrInvalidData, s)
	}

	startX, err := strconv.ParseInt(start[0], 10, 64)
	if err != nil {
		return Vents{}, fmt.Errorf("%w: invalid start X: %q",
			data.ErrInvalidData, s)
	}

	startY, err := strconv.ParseInt(start[1], 10, 64)
	if err != nil {
		return Vents{}, fmt.Errorf("%w: invalid start Y: %q",
			data.ErrInvalidData, s)
	}

	endX, err := strconv.ParseInt(end[0], 10, 64)
	if err != nil {
		return Vents{}, fmt.Errorf("%w: invalid start X: %q",
			data.ErrInvalidData, s)
	}

	endY, err := strconv.ParseInt(end[1], 10, 64)
	if err != nil {
		return Vents{}, fmt.Errorf("%w: invalid start Y: %q",
			data.ErrInvalidData, s)
	}

	return Vents{
		StartX: startX,
		StartY: startY,
		EndX:   endX,
		EndY:   endY,
	}, nil
}
