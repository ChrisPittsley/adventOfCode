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
	var last, increase int
	for line, text := range strings.Split(string(data), "\n") {
		if text == "" {
			continue
		}
		current, err := strconv.Atoi(text)
		if err != nil {
			lib.ErrorOut(fmt.Errorf("error on line %d: %v", line, err))
		}
		if line == 0 {
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
