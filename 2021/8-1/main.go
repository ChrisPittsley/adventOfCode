package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	count := 0
	for l, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " | ")
		if len(parts) != 2 {
			lib.ErrorOut(fmt.Errorf("error on line %d: bad syntax: %s", l, line))
		}
		for _, digit := range strings.Split(parts[1], " ") {
			switch len(digit) {
			case 2, 3, 4, 7:
				count += 1
			}
		}
	}
	fmt.Printf("%d\n", count)
	os.Exit(0)
}
