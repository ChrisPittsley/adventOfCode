package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
)

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	var floor, pos int
	for i, char := range data {
		switch char {
		case 40:
			floor += 1
		case 41:
			floor -= 1
		default:
			lib.ErrorOut(fmt.Errorf("bad input at index %d: '%s'", i, string(char)))
		}
		if floor == -1 {
			pos = i + 1
			break
		}
	}
	fmt.Printf("%d\n", pos)
	os.Exit(0)
}
