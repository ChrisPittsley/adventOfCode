package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func consumption(start, end int) int {
	if start > end {
		start, end = end, start
	}
	distance := end - start
	return (distance * (distance + 1)) / 2
}

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	var buckets = make(map[int]int)
	for p, n := range strings.Split(strings.TrimSuffix(string(data), "\n"), ",") {
		x, err := strconv.Atoi(n)
		if err != nil {
			lib.ErrorOut(fmt.Errorf("error at position %d: %v", p, err))
		}
		if _, ok := buckets[x]; !ok {
			buckets[x] = 0
		}
		buckets[x] += 1
	}
	min := 0
	for endPos := range buckets {
		fuel := 0
		for startPos, count := range buckets {
			fuel += consumption(startPos, endPos) * count
		}
		if fuel < min || min == 0 {
			min = fuel
		}
	}
	fmt.Printf("%d\n", min)
	os.Exit(0)
}
