package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
)

type coords struct {
	x, y int
}

type octopus struct {
	level   int
	flashed bool
}

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	var cavern = make(map[coords]octopus)
	var pos coords
	for _, b := range data {
		switch b {
		case 10:
			pos.x = 0
			pos.y += 1
			continue
		case 48:
			cavern[pos] = octopus{level: 0}
		case 49:
			cavern[pos] = octopus{level: 1}
		case 50:
			cavern[pos] = octopus{level: 2}
		case 51:
			cavern[pos] = octopus{level: 3}
		case 52:
			cavern[pos] = octopus{level: 4}
		case 53:
			cavern[pos] = octopus{level: 5}
		case 54:
			cavern[pos] = octopus{level: 6}
		case 55:
			cavern[pos] = octopus{level: 7}
		case 56:
			cavern[pos] = octopus{level: 8}
		case 57:
			cavern[pos] = octopus{level: 9}
		}
		pos.x += 1
	}
	cycle := 0
	for ; ; cycle += 1 {
		var flashes int
		for pos, o := range cavern {
			if o.flashed {
				o.level = 0
				o.flashed = false
				flashes += 1
			}
			o.level += 1
			cavern[pos] = o
		}
		if flashes == len(cavern) {
			break
		}
		checkFlashes(cavern)
	}
	fmt.Printf("%d\n", cycle)
	os.Exit(0)
}

func checkFlashes(cavern map[coords]octopus) {
	var f bool
	for pos, o := range cavern {
		if o.level > 9 && !o.flashed {
			f = true
			flash(pos, cavern)
		}
	}
	if f {
		checkFlashes(cavern)
	}
}

func flash(pos coords, cavern map[coords]octopus) {
	o := cavern[pos]
	o.flashed = true
	cavern[pos] = o
	var splash coords
	for splash.y = pos.y - 1; splash.y <= pos.y+1; splash.y += 1 {
		for splash.x = pos.x - 1; splash.x <= pos.x+1; splash.x += 1 {
			oct, ok := cavern[splash]
			if ok {
				oct.level += 1
				cavern[splash] = oct
			}
		}
	}
}
