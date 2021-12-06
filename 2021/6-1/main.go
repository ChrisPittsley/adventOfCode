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
	var fishes []int
	for _, n := range strings.Split(string(data), ",") {
		fish, err := strconv.Atoi(strings.TrimSuffix(n, "\n"))
		if err != nil {
			lib.ErrorOut(err)
		}
		fishes = append(fishes, fish)
	}
	for day := 1; day <= 80; day += 1 {
		var spawn []int
		for f := range fishes {
			if fishes[f] == 0 {
				spawn = append(spawn, 8)
				fishes[f] = 6
			} else {
				fishes[f] -= 1
			}
		}
		fishes = append(fishes, spawn...)
	}
	fmt.Printf("%d\n", len(fishes))
	os.Exit(0)
}
