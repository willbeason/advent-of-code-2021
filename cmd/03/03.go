package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/willbeason/advent-of-code-2021/pkg/data"
)

func main() {
	lines, err := data.ReadLines("cmd/03/data.txt")
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

	nBits := len(lines[0])
	nLines := len(lines)
	counts := make([]int, nBits)

	for _, r := range rows {
		for i, b := range r.bits {
			if b {
				counts[i]++
			}
		}
	}

	var gamma, epsilon int

	fmt.Println(counts)
	for i, count := range counts {
		power := nBits - 1 - i
		if count > (nLines / 2) {
			gamma += 1 << power
		} else {
			epsilon += 1 << power
		}
	}

	fmt.Println(gamma * epsilon)

	o2s := rows

	for i := 0; i < nBits; i++ {
		o2s = filterRows(o2s, i, true)
		if len(o2s) == 1 {
			break
		}
	}

	if len(o2s) > 1 {
		fmt.Printf("got multiple results: %+v\n", o2s)
		os.Exit(1)
	}

	o2 := 0

	for i, b := range o2s[0].bits {
		power := nBits - 1 - i
		if b {
			o2 += 1 << power
		}
	}

	co2s := rows
	for i := 0; i < nBits; i++ {
		co2s = filterRows(co2s, i, false)
		if len(co2s) == 1 {
			break
		}
	}

	if len(co2s) > 1 {
		fmt.Printf("got multiple results: %+v\n", co2s)
		os.Exit(1)
	}

	co2 := 0

	for i, b := range co2s[0].bits {
		power := nBits - 1 - i
		if b {
			co2 += 1 << power
		}
	}

	fmt.Println(o2, co2)
	fmt.Println(o2 * co2)
}

type row struct {
	bits []bool
}

func (r row) String() string {
	result := strings.Builder{}

	for _, b := range r.bits {
		if b {
			result.WriteString("1")
		} else {
			result.WriteString("0")
		}
	}

	return result.String()
}

func read(line string) (row, error) {
	result := row{
		bits: make([]bool, len(line)),
	}

	for i, c := range line {
		switch c {
		case '0':
			result.bits[i] = false
		case '1':
			result.bits[i] = true
		default:
			return row{}, fmt.Errorf("%w: %q",
				data.ErrInvalidData, line)
		}
	}

	return result, nil
}

func filterRows(rows []row, idx int, mostCommon bool) []row {
	idxCount := 0

	for _, r := range rows {
		if r.bits[idx] {
			idxCount++
		}
	}

	nRows := len(rows)
	result := make([]row, 0, nRows/2)

	var want bool
	if mostCommon {
		want = idxCount >= (nRows - idxCount)
	} else {
		want = idxCount < (nRows - idxCount)
	}

	for _, r := range rows {
		if r.bits[idx] == want {
			result = append(result, r)
		}
	}

	return result
}
