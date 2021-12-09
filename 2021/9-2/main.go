package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
)

type coords struct {
	x, y int
}

type point struct {
	height uint8
	mapped bool
}

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	heightMap := make(map[coords]point)
	row, col := 0, 0
	for pos, char := range data {
		p := coords{x: row, y: col}
		switch char {
		case 48:
			heightMap[p] = point{height: 0}
		case 49:
			heightMap[p] = point{height: 1}
		case 50:
			heightMap[p] = point{height: 2}
		case 51:
			heightMap[p] = point{height: 3}
		case 52:
			heightMap[p] = point{height: 4}
		case 53:
			heightMap[p] = point{height: 5}
		case 54:
			heightMap[p] = point{height: 6}
		case 55:
			heightMap[p] = point{height: 7}
		case 56:
			heightMap[p] = point{height: 8}
		case 57:
			heightMap[p] = point{height: 9}
		case 10:
			if pos != len(data)-1 {
				col = 0
				row += 1
				continue
			}
		default:
			lib.ErrorOut(fmt.Errorf("error at byte %d: bad input '%s'", pos, string(char)))
		}
		col += 1
	}
	var basins [3]int
	for xy, p1 := range heightMap {
		if p2, ok := heightMap[coords{x: xy.x, y: xy.y - 1}]; ok {
			if p1.height >= p2.height {
				continue
			}
		}
		if p2, ok := heightMap[coords{x: xy.x, y: xy.y + 1}]; ok {
			if p1.height >= p2.height {
				continue
			}
		}
		if p2, ok := heightMap[coords{x: xy.x - 1, y: xy.y}]; ok {
			if p1.height >= p2.height {
				continue
			}
		}
		if p2, ok := heightMap[coords{x: xy.x + 1, y: xy.y}]; ok {
			if p1.height >= p2.height {
				continue
			}
		}
		basin := mapBasin(xy, heightMap)
		switch {
		case basin > basins[0]:
			basins[0], basins[1], basins[2] = basin, basins[0], basins[1]
		case basin > basins[1]:
			basins[1], basins[2] = basin, basins[1]
		case basin > basins[2]:
			basins[2] = basin
		}
	}
	fmt.Printf("%d\n", basins[0]*basins[1]*basins[2])
	os.Exit(0)
}

func mapBasin(c coords, m map[coords]point) int {
	p, ok := m[c]
	if !ok || p.mapped || p.height == 9 {
		return 0
	}
	p.mapped = true
	m[c] = p
	size := 1
	for _, adj := range [4]coords{{x: c.x, y: c.y + 1}, {x: c.x, y: c.y - 1}, {x: c.x + 1, y: c.y}, {x: c.x - 1, y: c.y}} {
		size += mapBasin(adj, m)
	}
	return size
}
