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
	floor := 0
	for i, char := range data {
		switch char {
		case 40:
			floor += 1
		case 41:
			floor -= 1
		default:
			lib.ErrorOut(fmt.Errorf("bad input at index %d: '%s'", i, string(char)))
		}
	}
	fmt.Printf("%d\n", floor)
	os.Exit(0)
}
