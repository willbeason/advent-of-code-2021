package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/willbeason/advent-of-code-2021/pkg/data"
)

func main() {
	lines, err := data.ReadLines("cmd/04/data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	numbers, err := readNumbers(lines[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	boards, err := readBoards(lines[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

outer:
	for _, n := range numbers {
		fmt.Println(n)
		for i, b := range boards {
			score, won := b.Mark(n)
			if won {
				fmt.Println(i, score)
				break outer
			}
		}
	}

	boards, err = readBoards(lines[2:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	nWon := 0
	alreadyWon := make([]bool, len(boards))

outer2:
	for _, n := range numbers {
		fmt.Println(n)
		for i, b := range boards {
			if alreadyWon[i] {
				continue
			}

			score, won := b.Mark(n)
			if won {
				alreadyWon[i] = true
				nWon++

				if nWon == len(boards) {
					fmt.Println(score)
					break outer2
				}
			}
		}
	}
}

func readNumbers(s string) ([]int64, error) {
	ss := strings.Split(s, ",")
	result := make([]int64, len(ss))

	var err error
	for i, n := range ss {
		result[i], err = strconv.ParseInt(n, 10, 64)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func readBoards(lines []string) ([]*board, error) {
	var boards []*board

	for i := 0; i < len(lines); i += 6 {
		b, err := readBoard(lines[i : i+5])
		if err != nil {
			return nil, err
		}

		boards = append(boards, &b)
	}

	return boards, nil
}

func readBoard(lines []string) (board, error) {
	b := board{
		numbers: make([]int64, 25),
		marked:  make([]bool, 25),
	}

	for i, line := range lines {
		ns, err := readLine(line)
		if err != nil {
			return board{}, nil
		}

		for j, n := range ns {
			idx := i*5 + j

			b.numbers[idx] = n
		}
	}

	return b, nil
}

func readLine(line string) ([]int64, error) {
	var result []int64

	parts := strings.Split(line, " ")
	for _, part := range parts {
		if len(part) == 0 {
			continue
		}

		n, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}

	if len(result) != 5 {
		return nil, fmt.Errorf("%w: want 5 numbers: %q",
			data.ErrInvalidData, line)
	}

	return result, nil
}

type board struct {
	numbers []int64
	marked  []bool
}

func (b board) Mark(i int64) (int64, bool) {
	marked := -1
	for idx, n := range b.numbers {
		if n == i {
			b.marked[idx] = true
			marked = idx
			break
		}
	}

	if marked == -1 {
		return 0, false
	}

	row := marked / 5
	if b.rowWin(row) {
		return b.score(i), true
	}

	column := marked % 5
	if b.columnWin(column) {
		return b.score(i), true
	}

	return 0, false
}

func (b board) rowWin(i int) bool {
	for j := i * 5; j < (i+1)*5; j++ {
		if !b.marked[j] {
			return false
		}
	}

	return true
}

func (b board) columnWin(i int) bool {
	for j := i; j < 25; j += 5 {
		if !b.marked[j] {
			return false
		}
	}

	return true
}

func (b board) score(i int64) int64 {
	var score int64

	for idx, n := range b.numbers {
		if !b.marked[idx] {
			score += n
		}
	}

	return score * i
}
