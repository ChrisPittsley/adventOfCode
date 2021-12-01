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
	var grid [1000][1000]bool
	for l, line := range strings.Split(string(data), "\n") {
		if len(line) == 0 {
			continue
		}
		tokens := strings.Split(line, " ")
		switch tokens[0] {
		case "turn":
			if len(tokens) != 5 {
				lib.ErrorOut(fmt.Errorf("bad input on line %d: 'turn' takes 4 args", l))
			}
			x, y, err := parseCoords(tokens[2], tokens[4])
			if err != nil {
				lib.ErrorOut(fmt.Errorf("bad input on line %d: %v", l, err))
			}
			switch tokens[1] {
			case "on":
				on(&grid, x, y)
			case "off":
				off(&grid, x, y)
			default:
				lib.ErrorOut(fmt.Errorf("bad input on line %d: bad argument to 'turn': '%s'", l, tokens[1]))
			}
		case "toggle":
			if len(tokens) != 4 {
				lib.ErrorOut(fmt.Errorf("bad input on line %d: 'turn' takes 3 args", l))
			}
			x, y, err := parseCoords(tokens[1], tokens[3])
			if err != nil {
				lib.ErrorOut(fmt.Errorf("bad input on line %d: %v", l, err))
			}
			toggle(&grid, x, y)
		}
	}
	fmt.Printf("%d\n", total(grid))
	os.Exit(0)
}

func parseCoords(startPoint, endPoint string) ([2]int, [2]int, error) {
	var x, y [2]int
	var err error
	start := strings.Split(startPoint, ",")
	if len(start) != 2 {
		return x,y, fmt.Errorf("bad start coordinates: '%s'", startPoint)
	}
	end := strings.Split(endPoint, ",")
	if len(end) != 2 {
		return x,y, fmt.Errorf("bad end coordinates: '%s'", endPoint)
	}
	x[0], err = strconv.Atoi(start[0])
	if err != nil {
		return x, y, fmt.Errorf("bad start coordinates: '%s': %v", startPoint, err)
	}
	x[1], err = strconv.Atoi(end[0])
	if err != nil {
		return x, y, fmt.Errorf("bad end coordinates: '%s': %v", endPoint, err)
	}
	y[0], err = strconv.Atoi(start[1])
	if err != nil {
		return x, y, fmt.Errorf("bad start coordinates: '%s': %v", startPoint, err)
	}
	y[1], err = strconv.Atoi(end[1])
	if err != nil {
		return x, y, fmt.Errorf("bad end coordinates: '%s': %v", endPoint, err)
	}
	return x, y, nil
}

func on(grid *[1000][1000]bool, xx, yy [2]int) {
	for x := xx[0]; x <= xx[1]; x += 1 {
		for y := yy[0]; y <= yy[1]; y += 1 {
			grid[x][y] = true
		}
	}
}

func off(grid *[1000][1000]bool, xx, yy [2]int) {
	for x := xx[0]; x <= xx[1]; x += 1 {
		for y := yy[0]; y <= yy[1]; y += 1 {
			grid[x][y] = false
		}
	}
}

func toggle(grid *[1000][1000]bool, xx, yy [2]int) {
	for x := xx[0]; x <= xx[1]; x += 1 {
		for y := yy[0]; y <= yy[1]; y += 1 {
			grid[x][y] = !grid[x][y]
		}
	}
}

func total(grid [1000][1000]bool) int {
	t := 0
	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] {
				t += 1
			}
		}
	}
	return t
}