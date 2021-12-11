package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/willbeason/advent-of-code-2021/pkg/data"
)

func main() {
	lines, err := data.ReadLines("cmd/10/data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	corruptionScore := 0

	for _, line := range lines {
		ls := CodeLine(line).CorruptionScore()
		corruptionScore += ls
	}

	fmt.Println(corruptionScore)

	var completionScores []int

	for _, line := range lines {
		ls := CodeLine(line).CompletionScore()
		if ls != 0 {
			completionScores = append(completionScores, ls)
		}
	}

	sort.Ints(completionScores)

	middle := len(completionScores) / 2

	fmt.Println(completionScores[middle])
}

type CodeLine string

func (cl CodeLine) CorruptionScore() int {
	want := ""

	for _, c := range cl {
		switch c {
		case '(':
			want = ")" + want
		case '[':
			want = "]" + want
		case '{':
			want = "}" + want
		case '<':
			want = ">" + want
		case ')':
			if want[0] != ')' {
				return 3
			}

			want = want[1:]
		case ']':
			if want[0] != ']' {
				return 57
			}

			want = want[1:]
		case '}':
			if want[0] != '}' {
				return 1197
			}

			want = want[1:]
		case '>':
			if want[0] != '>' {
				return 25137
			}

			want = want[1:]
		}
	}

	return 0
}

func (cl CodeLine) CompletionScore() int {
	want := ""

	for _, c := range cl {
		switch c {
		case '(':
			want = ")" + want
		case '[':
			want = "]" + want
		case '{':
			want = "}" + want
		case '<':
			want = ">" + want
		case ')':
			if want[0] != ')' {
				return 0
			}

			want = want[1:]
		case ']':
			if want[0] != ']' {
				return 0
			}

			want = want[1:]
		case '}':
			if want[0] != '}' {
				return 0
			}

			want = want[1:]
		case '>':
			if want[0] != '>' {
				return 0
			}

			want = want[1:]
		}
	}

	score := 0
	for _, c := range want {
		score *= 5

		switch c {
		case ')':
			score++
		case ']':
			score += 2
		case '}':
			score += 3
		case '>':
			score += 4
		}
	}

	return score
}
