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
	var x, y [2]int
	var santa int
	for i, char := range data {
		switch char {
		case 60: x[santa] -= 1
		case 62: x[santa] += 1
		case 86: y[santa] -= 1
		case 94: y[santa] += 1
		case 118: y[santa] -= 1
		default:
			lib.ErrorOut(fmt.Errorf("bad input at position %d: '%s'", i, string(char)))
		}
		houses[fmt.Sprintf("%d,%d", x[santa], y[santa])] = true
		if santa == 0 {
			santa = 1
		} else {
			santa = 0
		}
	}
	fmt.Printf("%d\n", len(houses))
	os.Exit(0)
}
