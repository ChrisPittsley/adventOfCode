package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
	"sort"
)

const (
	NewLine byte = 10
	ParenL  byte = 40
	ParenR  byte = 41
	AngleL  byte = 60
	AngleR  byte = 62
	SquareL byte = 91
	SquareR byte = 93
	CurlyL  byte = 123
	CurlyR  byte = 125
)

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	var points = map[byte]int{ParenL: 1, SquareL: 2, CurlyL: 3, AngleL: 4}
	var pairs = map[byte]byte{ParenR: ParenL, SquareR: SquareL, CurlyR: CurlyL, AngleR: AngleL}
	var stack []byte
	var scores []int
	var skip bool
	for _, b := range data {
		if skip && b != NewLine {
			continue
		}
		head := len(stack) - 1
		switch b {
		case ParenL, AngleL, SquareL, CurlyL:
			stack = append(stack, b)
		case ParenR, AngleR, SquareR, CurlyR:
			if head < 0 {
				skip = true
				continue
			}
			if pairs[b] != stack[head] {
				skip = true
				continue
			}
			stack = stack[:head]
		case NewLine:
			if skip || len(stack) == 0 {
				skip = false
				stack = []byte{}
				continue
			}
			score := 0
			for ; head >= 0; head -= 1 {
				score = (score * 5) + points[stack[head]]
			}
			scores = append(scores, score)
			stack = []byte{}
		}
	}
	sort.Ints(scores)
	median := len(scores) / 2
	fmt.Printf("%d\n", scores[median])
	os.Exit(0)
}
