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

type instruction struct {
	axis   string
	offset int
}

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	var paper = make(map[point]bool)
	var steps []instruction
	for l, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "fold along ") {
			instr := strings.Split(strings.TrimPrefix(line, "fold along "), "=")
			if len(instr) != 2 {
				lib.ErrorOut(fmt.Errorf("bad syntax on line %d", l))
			}
			if instr[0] != "x" && instr[0] != "y" {
				lib.ErrorOut(fmt.Errorf("bad syntax on line %d: bad axis '%s'", l, instr[0]))
			}
			offset, err := strconv.Atoi(instr[1])
			if err != nil {
				lib.ErrorOut(fmt.Errorf("bad syntax on line %d: bad offset '%s': %v", l, instr[1], err))
			}
			steps = append(steps, instruction{
				axis:   instr[0],
				offset: offset,
			})
			continue
		}
		coords := strings.Split(line, ",")
		if len(coords) != 2 {
			lib.ErrorOut(fmt.Errorf("bad syntax on line %d", l))
		}
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			lib.ErrorOut(fmt.Errorf("bad syntax on line %d: bad x coordinate '%s': %v", l, coords[0], err))
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			lib.ErrorOut(fmt.Errorf("bad syntax on line %d: bad y coordinate '%s': %v", l, coords[1], err))
		}
		paper[point{x: x, y: y}] = true
	}
	for _, step := range steps {
		fold(paper, step)
	}
	fmt.Printf("%s\n", draw(paper))
	os.Exit(0)
}

func fold(paper map[point]bool, i instruction) {
	switch i.axis {
	case "x":
		for p := range paper {
			move := (i.offset - p.x) * 2
			if move > 0 {
				continue
			}
			delete(paper, p)
			p.x = p.x + move
			paper[p] = true
		}
	case "y":
		for p := range paper {
			move := (i.offset - p.y) * 2
			if move > 0 {
				continue
			}
			delete(paper, p)
			p.y = p.y + move
			paper[p] = true
		}
	}
}

func draw(paper map[point]bool) string {
	corner := point{}
	for p := range paper {
		if p.x > corner.x {
			corner.x = p.x
		}
		if p.y > corner.y {
			corner.y = p.y
		}
	}
	var out string
	for y := 0; y <= corner.y; y += 1 {
		for x := 0; x <= corner.x; x += 1 {
			_, ok := paper[point{x: x, y: y}]
			if !ok {
				out += " "
			} else {
				out += "#"
			}
		}
		out += "\n"
	}
	return out
}
