package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
)

type point struct {
	row, col int
}

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	heightMap := make(map[point]uint8)
	row, col := 0, 0
	for pos, char := range data {
		p := point{row: row, col: col}
		switch char {
		case 48:
			heightMap[p] = 0
		case 49:
			heightMap[p] = 1
		case 50:
			heightMap[p] = 2
		case 51:
			heightMap[p] = 3
		case 52:
			heightMap[p] = 4
		case 53:
			heightMap[p] = 5
		case 54:
			heightMap[p] = 6
		case 55:
			heightMap[p] = 7
		case 56:
			heightMap[p] = 8
		case 57:
			heightMap[p] = 9
		case 10:
			if pos != len(data)-1 {
				col = 0
				row += 1
				continue
			}
		default:
			lib.ErrorOut(fmt.Errorf("error at byte %d: bad input '%s'", pos, string(char)))
		}
		col += 1
	}
	var risk uint
	for p, h1 := range heightMap {
		if h2, ok := heightMap[point{row: p.row, col: p.col - 1}]; ok {
			if h1 >= h2 {
				continue
			}
		}
		if h2, ok := heightMap[point{row: p.row, col: p.col + 1}]; ok {
			if h1 >= h2 {
				continue
			}
		}
		if h2, ok := heightMap[point{row: p.row - 1, col: p.col}]; ok {
			if h1 >= h2 {
				continue
			}
		}
		if h2, ok := heightMap[point{row: p.row + 1, col: p.col}]; ok {
			if h1 >= h2 {
				continue
			}
		}
		risk += uint(h1) + 1
	}
	fmt.Printf("%d\n", risk)
	os.Exit(0)
}
