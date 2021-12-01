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
	sums := make([]int, len(lines), len(lines))
	for line, text := range lines {
		if text == "" {
			continue
		}
		num, err := strconv.Atoi(text)
		if err != nil {
			lib.ErrorOut(fmt.Errorf("error on line %d: %v", line, err))
		}
		if line < len(lines) - 2 {
			sums[line] += num
		}
		if line > 2 {
			sums[line - 2] += num
		}
		if line > 1 {
			sums[line - 1] += num
		}
	}
	var last, increase int
	for n, current := range sums {
		if n == 0 {
			last = current
			continue
		}
		if last < current {
			increase += 1
		}
		last = current
	}
	fmt.Printf("%d\n", increase)
	os.Exit(0)
}
