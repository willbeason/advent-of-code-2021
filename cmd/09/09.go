package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/willbeason/advent-of-code-2021/pkg/data"
)

func main() {
	lines, err := data.ReadLines("cmd/09/data.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	hm := readHeightMap(lines)

	score := hm.Score()
	fmt.Println(score)

	score2 := hm.Score2()
	fmt.Println(score2)
}

func readHeightMap(lines []string) HeightMap {
	hm := HeightMap{}

	for _, l := range lines {
		hLine := make([]int32, len(l))
		for i, c := range l {
			hLine[i] = c - '0'
		}

		hm.Heights = append(hm.Heights, hLine)
	}

	return hm
}

type HeightMap struct {
	Heights [][]int32
}

func (hm *HeightMap) IsLowest(i, j int) bool {
	hLine := hm.Heights[i]
	h := hLine[j]

	if j > 0 && hm.Heights[i][j-1] <= h {
		return false
	}

	if j < len(hLine)-1 && hm.Heights[i][j+1] <= h {
		return false
	}

	if i > 0 && hm.Heights[i-1][j] <= h {
		return false
	}

	if i < len(hm.Heights)-1 && hm.Heights[i+1][j] <= h {
		return false
	}

	return true
}

func (hm *HeightMap) Score() int32 {
	var score int32

	for i := 0; i < len(hm.Heights); i++ {
		hLine := hm.Heights[i]

		for j, h := range hLine {
			if hm.IsLowest(i, j) {
				score = score + 1 + h
			}
		}
	}

	return score
}

func (hm *HeightMap) LowestNeighbor(i, j int) (int, int) {
	hLine := hm.Heights[i]
	h := hLine[j]

	li, lj := i, j

	if j > 0 && hm.Heights[i][j-1] < h {
		h = hm.Heights[i][j-1]
		li, lj = i, j-1
	}

	if j < len(hLine)-1 && hm.Heights[i][j+1] < h {
		h = hm.Heights[i][j+1]
		li, lj = i, j+1
	}

	if i > 0 && hm.Heights[i-1][j] < h {
		h = hm.Heights[i-1][j]
		li, lj = i-1, j
	}

	if i < len(hm.Heights)-1 && hm.Heights[i+1][j] < h {
		h = hm.Heights[i+1][j]
		li, lj = i+1, j
	}

	return li, lj
}

func (hm *HeightMap) Score2() int {
	scores := make([][]int, len(hm.Heights))
	for i, iLine := range hm.Heights {
		scores[i] = make([]int, len(iLine))
	}

	for i := 0; i < len(hm.Heights); i++ {
		hLine := hm.Heights[i]

		for j, h := range hLine {
			if h == 9 {
				continue
			}

			scores[i][j]++

			pi, pj := i, j
			li, lj := hm.LowestNeighbor(i, j)

			for pi != li || pj != lj {
				scores[li][lj]++

				pi, pj = li, lj
				li, lj = hm.LowestNeighbor(li, lj)
			}
		}
	}

	var minScores []int

	for i, iScores := range scores {
		for j, score := range iScores {
			if hm.IsLowest(i, j) {
				minScores = append(minScores, score)
			}
		}
	}

	sort.Slice(minScores, func(i, j int) bool {
		return minScores[i] > minScores[j]
	})

	return minScores[0] * minScores[1] * minScores[2]
}
