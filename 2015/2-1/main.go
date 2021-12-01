package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	lines := strings.Split(string(data), "\n")
	var total int
	for i, line := range lines {
		if i+1 == len(lines) {
			break
		}
		dimensions := strings.Split(line, "x")
		if len(dimensions) != 3 {
			lib.ErrorOut(fmt.Errorf("bad input on line %d: '%s'", i, line))
		}
		var smallest int
		for edge1 := range dimensions {
			x, err := strconv.Atoi(dimensions[edge1])
			if err != nil {
				lib.ErrorOut(fmt.Errorf("bad input on line %d: '%s': %v", i, line, err))
			}
			for edge2 := range dimensions[edge1+1:] {
				y, err := strconv.Atoi(dimensions[edge1+edge2+1])
				if err != nil {
					lib.ErrorOut(fmt.Errorf("bad input on line %d: '%s': %v", i, line, err))
				}
				area := x * y
				total += area * 2
				switch {
				case smallest == 0:
					smallest = area
				case area < smallest:
					smallest = area
				}
			}
		}
		total += smallest
	}
	fmt.Printf("%d\n", total)
	os.Exit(0)
}
