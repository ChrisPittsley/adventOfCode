package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type space struct {
	val    int
	marked bool
}

type board [5][5]space

func (b *board) score() int {
	var sum int
	for row := range b {
		for col := range b[row] {
			if !b[row][col].marked {
				sum += b[row][col].val
			}
		}
	}
	return sum
}

func (b *board) mark(num int) int {
	var win bool
	var colWin [5]bool
	for col := range colWin {
		colWin[col] = true
	}
	for row := range b {
		var rowWin = true
		for col := range b[row] {
			if b[row][col].val == num {
				b[row][col].marked = true
			}
			rowWin = b[row][col].marked && rowWin
			colWin[col] = b[row][col].marked && colWin[col]
		}
		if rowWin {
			win = true
		}
	}
	for col := range colWin {
		if colWin[col] {
			win = true
		}
	}
	if win {
		return b.score()
	}
	return 0
}

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	var draws []int
	var boards []board
	lines := strings.Split(string(data), "\n")
	for l := 0; l < len(lines); {
		switch {
		case l == 0:
			draws, err = parseDraws(lines[l])
			if err != nil {
				lib.ErrorOut(fmt.Errorf("error on line %d: %v", l, err))
			}
			l += 1
		case lines[l] == "":
			l += 1
			continue
		default:
			b, err := parseBoard(lines[l : l+5])
			if err != nil {
				lib.ErrorOut(fmt.Errorf("error parsing board on line %d: %v", l, err))
			}
			boards = append(boards, b)
			l += 5
		}
	}
	var score int
Draw:
	for _, draw := range draws {
		var losers []board
		for b := range boards {
			score = boards[b].mark(draw) * draw
			if score > 0 {
				if len(boards) == 1 {
					break Draw
				}
			} else {
				losers = append(losers, boards[b])
			}
		}
		boards = losers
	}
	fmt.Printf("%d\n", score)
	os.Exit(0)
}

func parseDraws(in string) ([]int, error) {
	var draws []int
	for _, s := range strings.Split(in, ",") {
		if s == "" {
			continue
		}
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		draws = append(draws, n)
	}
	return draws, nil
}

func parseBoard(in []string) (board, error) {
	var b board
	if len(in) != 5 {
		return board{}, fmt.Errorf("input %v has %d rows, should have 5", in, len(in))
	}
	for r, rowIn := range in {
		var row []string
		for _, s := range strings.Split(rowIn, " ") {
			if s == "" {
				continue
			}
			row = append(row, s)
		}
		if len(row) != 5 {
			return board{}, fmt.Errorf("input %v has %d columns, should have 5", row, len(row))
		}
		for c, s := range row {
			n, err := strconv.Atoi(s)
			if err != nil {
				return board{}, err
			}
			b[r][c] = space{
				val: n,
			}
		}
	}
	return b, nil
}
