package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	grid := make(map[point]int)
	for l, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		points := strings.Split(line, " -> ")
		if len(points) != 2 {
			lib.ErrorOut(fmt.Errorf("bad syntax on line %d: %s", l, line))
		}
		p1, err := parsePoint(points[0])
		if err != nil {
			lib.ErrorOut(fmt.Errorf("bad syntax on line %d: %v", l, err))
		}
		p2, err := parsePoint(points[1])
		if err != nil {
			lib.ErrorOut(fmt.Errorf("bad syntax on line %d: %v", l, err))
		}
		switch {
		case p1.x == p2.x:
			if p1.y > p2.y {
				p1, p2 = p2, p1
			}
			for p1.y <= p2.y {
				if _, ok := grid[p1]; !ok {
					grid[p1] = 1
				} else {
					grid[p1] += 1
				}
				p1.y += 1
			}
		case p1.y == p2.y:
			if p1.x > p2.x {
				p1, p2 = p2, p1
			}
			for p1.x <= p2.x {
				if _, ok := grid[p1]; !ok {
					grid[p1] = 1
				} else {
					grid[p1] += 1
				}
				p1.x += 1
			}
		default:
			continue
		}
	}
	var hotPoints int
	for _, p := range grid {
		if p >= 2 {
			hotPoints += 1
		}
	}
	fmt.Printf("%d\n", hotPoints)
	os.Exit(0)
}

func parsePoint(in string) (point, error) {
	rawCoords := strings.Split(in, ",")
	if len(rawCoords) != 2 {
		return point{}, fmt.Errorf("bad coordinates: %s", in)
	}
	x, err := strconv.Atoi(rawCoords[0])
	if err != nil {
		return point{}, fmt.Errorf("bad coordinates: %s: %v", in, err)
	}
	y, err := strconv.Atoi(rawCoords[1])
	if err != nil {
		return point{}, fmt.Errorf("bad coordinates: %s: %v", in, err)
	}
	return point{x: x, y: y}, nil
}
