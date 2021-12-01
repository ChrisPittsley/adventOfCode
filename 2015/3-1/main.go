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
	houses := map[string]bool{"0,0":true}
	var x, y int
	for i, char := range data {
		switch char {
		case 60: x -= 1
		case 62: x += 1
		case 86: y -= 1
		case 94: y += 1
		case 118: y -= 1
		default:
			lib.ErrorOut(fmt.Errorf("bad input at position %d: '%s'", i, string(char)))
		}
		houses[fmt.Sprintf("%d,%d", x, y)] = true
	}
	fmt.Printf("%d\n", len(houses))
	os.Exit(0)
}
