package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/willbeason/advent-of-code-2021/pkg/data"
)

func main() {
	lines, err := data.ReadLines("cmd/08/data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	easyDigits := 0

	for _, line := range lines {
		parts := strings.Split(line, " | ")
		if len(parts) != 2 {
			panic(line)
		}

		decoder := makeDecoder(strings.Split(parts[0], " "))

		for _, digit := range strings.Split(parts[1], " ") {
			result := decoder.Decode(digit)
			switch result {
			case 1, 4, 7, 8:
				easyDigits++
			}
		}
	}

	fmt.Println(easyDigits)

	sum := 0

	for _, line := range lines {
		parts := strings.Split(line, " | ")
		if len(parts) != 2 {
			panic(line)
		}

		decoder := makeDecoder(strings.Split(parts[0], " "))

		result := decoder.Decode(parts[1])
		sum += result
	}

	fmt.Println(sum)
}

func makeDecoder(patterns []string) Decoder {
	var d0, d1, d2, d3, d4, d5, d6, d7, d8, d9 [7]bool

	var d235 [][7]bool // 5
	var d069 [][7]bool // 6

	for _, pattern := range patterns {
		ons := toOn(pattern)

		switch count(ons) {
		case 2:
			d1 = ons
		case 3:
			d7 = ons
		case 4:
			d4 = ons
		case 5:
			d235 = append(d235, ons)
		case 6:
			d069 = append(d069, ons)
		case 7:
			d8 = ons
		default:
			panic(fmt.Sprintf("bad decoder pattern %q", pattern))
		}
	}

	switch {
	case intersectCount(d235[0], d235[1]) == 3:
		d3 = d235[2]
		if intersectCount(d235[0], d4) == 2 {
			d2 = d235[0]
			d5 = d235[1]
		} else {
			d2 = d235[1]
			d5 = d235[0]
		}
	case intersectCount(d235[0], d235[2]) == 3:
		if intersectCount(d235[0], d4) == 2 {
			d2 = d235[0]
			d5 = d235[2]
		} else {
			d2 = d235[2]
			d5 = d235[0]
		}
		d3 = d235[1]
	case intersectCount(d235[1], d235[2]) == 3:
		d3 = d235[0]
		if intersectCount(d235[1], d4) == 2 {
			d2 = d235[1]
			d5 = d235[2]
		} else {
			d2 = d235[2]
			d5 = d235[1]
		}
	}

	switch {
	case intersectCount(d069[0], d1) == 1:
		d6 = d069[0]
		if intersectCount(d069[1], d5) == 5 {
			d9 = d069[1]
			d0 = d069[2]
		} else {
			d9 = d069[2]
			d0 = d069[1]
		}
	case intersectCount(d069[1], d1) == 1:
		d6 = d069[1]
		if intersectCount(d069[0], d5) == 5 {
			d9 = d069[0]
			d0 = d069[2]
		} else {
			d9 = d069[2]
			d0 = d069[0]
		}
	case intersectCount(d069[2], d1) == 1:
		d6 = d069[2]
		if intersectCount(d069[0], d5) == 5 {
			d9 = d069[0]
			d0 = d069[1]
		} else {
			d9 = d069[1]
			d0 = d069[0]
		}
	}

	return Decoder{
		d0: 0,
		d1: 1,
		d2: 2,
		d3: 3,
		d4: 4,
		d5: 5,
		d6: 6,
		d7: 7,
		d8: 8,
		d9: 9,
	}
}

func toOn(pattern string) [7]bool {
	var result [7]bool

	for _, c := range pattern {
		switch c {
		case 'a':
			result[0] = true
		case 'b':
			result[1] = true
		case 'c':
			result[2] = true
		case 'd':
			result[3] = true
		case 'e':
			result[4] = true
		case 'f':
			result[5] = true
		case 'g':
			result[6] = true
		default:
			panic(fmt.Sprintf("bad on pattern %q", pattern))
		}
	}

	return result
}

func intersectCount(b1, b2 [7]bool) int {
	result := 0
	for i := range b1 {
		if b1[i] && b2[i] {
			result++
		}
	}

	return result
}

func count(b [7]bool) int {
	result := 0
	for _, t := range b {
		if t {
			result++
		}
	}

	return result
}

type Decoder map[[7]bool]int

func (d Decoder) Decode(s string) int {
	parts := strings.Split(s, " ")

	result := 0
	for _, p := range parts {
		result *= 10

		ons := toOn(p)
		digit, found := d[ons]
		if !found {
			panic(fmt.Sprintf("bad decode pattern %q", s))
		}

		result += digit
	}

	return result
}
