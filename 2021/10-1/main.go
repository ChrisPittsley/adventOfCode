package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
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
	var points = map[byte]int{ParenR: 3, SquareR: 57, CurlyR: 1197, AngleR: 25137}
	var pairs = map[byte]byte{ParenR: ParenL, SquareR: SquareL, CurlyR: CurlyL, AngleR: AngleL}
	var stack []byte
	var score int
	var skip bool
	for _, b := range data {
		if skip && b != NewLine {
			continue
		}
		skip = false
		head := len(stack) - 1
		switch b {
		case ParenL, AngleL, SquareL, CurlyL:
			stack = append(stack, b)
		case ParenR, AngleR, SquareR, CurlyR:
			if head < 0 {
				score += points[b]
				skip = true
				continue
			}
			if pairs[b] != stack[head] {
				score += points[b]
				skip = true
				continue
			}
			stack = stack[:head]
		case NewLine:
			stack = []byte{}
		}
	}
	fmt.Printf("%d\n", score)
	os.Exit(0)
}
