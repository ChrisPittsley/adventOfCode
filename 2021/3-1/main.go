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
	var bits [12]int
	var bit, lines int
	var epsilon, gamma uint16
	for pos, char := range data {
		switch char {
		case 48:
			bit += 1
		case 49:
			bits[bit] += 1
			bit += 1
		case 10:
			lines += 1
			bit = 0
		default:
			lib.ErrorOut(fmt.Errorf("error at byte %d: invalid character", pos))
		}
	}
	for bit, sum := range bits {
		if lines-sum < sum {
			gamma += 2048 >> bit
		}
	}
	epsilon = gamma ^ 0x0FFF
	fmt.Printf("%d\n", uint32(gamma)*uint32(epsilon))
	os.Exit(0)
}
