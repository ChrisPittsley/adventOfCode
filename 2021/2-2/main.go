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
	var depth, hpos, aim int
	for line, text := range strings.Split(string(data), "\n") {
		if text == "" {
			continue
		}
		parsed := strings.Split(text, " ")
		if len(parsed) != 2 {
			lib.ErrorOut(fmt.Errorf("error on line %d: bad syntax", line))
		}
		val, err := strconv.Atoi(parsed[1])
		if err != nil {
			lib.ErrorOut(fmt.Errorf("error on line %d: %v", line, err))
		}
		switch parsed[0] {
		case "forward":
			hpos += val
			depth += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		default:
			lib.ErrorOut(fmt.Errorf("error on line %d: bad syntax", line))
		}
	}
	fmt.Printf("%d\n", depth*hpos)
	os.Exit(0)
}
