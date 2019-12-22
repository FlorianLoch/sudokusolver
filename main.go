package main

import (
	"fmt"
	"strings"
	"strconv"
)

var exampleBoard = Board{5, 3, 0, 0, 7, 0, 0, 0, 0, 6, 0, 0, 1, 9, 5, 0, 0, 0, 0, 9, 8, 0, 0, 0, 0, 6, 0, 8, 0, 0, 0, 6, 0, 0, 0, 3, 4, 0, 0, 8, 0, 3, 0, 0, 1, 7, 0, 0, 0, 2, 0, 0, 0, 6, 0, 6, 0, 0, 0, 0, 2, 8, 0, 0, 0, 0, 4, 1, 9, 0, 0, 5, 0, 0, 0, 0, 8, 0, 0, 7, 9}

/*
	Given situation:
	5, 3, 0, 0, 7, 0, 0, 0, 0,
	6, 0, 0, 1, 9, 5, 0, 0, 0,
	0, 9, 8, 0, 0, 0, 0, 6, 0,
	8, 0, 0, 0, 6, 0, 0, 0, 3,
	4, 0, 0, 8, 0, 3, 0, 0, 1,
	7, 0, 0, 0, 2, 0, 0, 0, 6,
	0, 6, 0, 0, 0, 0, 2, 8, 0,
	0, 0, 0, 4, 1, 9, 0, 0, 5,
	0, 0, 0, 0, 8, 0, 0, 7, 9

	Solved board:
	5, 3, 4, 6, 7, 8, 9, 1, 2,
	6, 7, 2, 1, 9, 5, 3, 4, 8,
	1, 9, 8, 3, 4, 2, 5, 6, 7,
	8, 5, 9, 7, 6, 1, 4, 2, 3,
	4, 2, 6, 8, 5, 3, 7, 9, 1,
	7, 1, 3, 9, 2, 4, 8, 5, 6,
	9, 6, 1, 5, 3, 7, 2, 8, 4,
	2, 8, 7, 4, 1, 9, 6, 3, 5,
	3, 4, 5, 2, 8, 6, 1, 7, 9
*/

var (
	blockOffsets = []int{0, 1, 2, 9, 10, 11, 18, 19, 20}
	opsCnt = 0
)

type Board []int

func NewBoard() Board {
	return make([]int, 81)
}

func (b Board) Set(x, y, v int) {
	idx := x + y * 9

	if idx < 0 || idx > 80 {
		panic(fmt.Sprintf("Cannot set field with index %d", idx))
	}

	b[idx] = v
}

func (b Board) Get(x, y int) int {
	idx := x + y * 9

	if idx < 0 || idx > 80 {
		panic(fmt.Sprintf("Cannot get field with index %d", idx))
	}

	return b[idx]
}

func (b Board) valuePossibleAt(v, idx int) bool {
	opsCnt++

	// Check row
	startOfRowIdx := idx / 9 * 9
	for i := 0; i < 9; i++ {
		if b[startOfRowIdx + i] == v {
			return false
		}
	}

	// Check column
	column := idx % 9
	for i := 0; i < 9; i++ {
		if b[column + i * 9] == v {
			return false
		}
	}

	// Check block
	row := idx / 9
	topLeftIdx := row / 3 * 3 * 9 + column / 3 * 3

	for _, offset := range blockOffsets {
		if b[topLeftIdx + offset] == v {
			return false
		}
	}

	return true
}

func (b Board) String() string {
	var sb strings.Builder

	for i := 0; i < 81; i++ {
		if i % 9 == 0 {
			sb.WriteString("\n")
		}
		sb.WriteString(strconv.Itoa(b[i]))
		sb.WriteRune(' ')
	}

	sb.WriteRune('\n')

	return sb.String()
}

func (b Board) Solve() {
	opsCnt = 0
	fmt.Printf("Trying to find solution for: %s\n", b)

	if b.solveInner(0) {
		fmt.Printf("Found valid solution after %d steps: %s", opsCnt, b)
	} else {
		fmt.Printf("Could not find a valid solution for this puzzle after %d steps!\n", opsCnt)
	}
}

func (b Board) solveInner(idx int) bool {
	// fmt.Printf("%d: %s\n", idx, b)

	// Find next field not being set already
	for idx < 81 && b[idx] != 0 {
		idx = idx + 1
	}

	if idx == 81 {
		// All fields are set, we found a solution for this puzzle
		return true
	}

	// Iterate over all possible values for this field
	for i := 1; i < 10; i++ {
		if b.valuePossibleAt(i, idx) {
			b[idx] = i
			if b.solveInner(idx + 1) {
				return true
			}
		}
	}

	// Reset field for backtracking
	b[idx] = 0

	return false
}

func main() {
	exampleBoard.Solve()
}