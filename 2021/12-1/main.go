package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
	"strings"
)

type node struct {
	neighbors []string
	small     bool
}

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	var cavern = map[string]node{"start": {}, "end": {}}
	for l, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		path := strings.Split(line, "-")
		if len(path) != 2 {
			lib.ErrorOut(fmt.Errorf("bad syntax on line %d", l))
		}
		n0, ok := cavern[path[0]]
		if !ok {
			n0 = node{small: isSmall(path[0])}
		}
		n1, ok := cavern[path[1]]
		if !ok {
			n1 = node{small: isSmall(path[1])}
		}
		n0.neighbors = append(n0.neighbors, path[1])
		n1.neighbors = append(n1.neighbors, path[0])
		cavern[path[0]] = n0
		cavern[path[1]] = n1
	}
	paths := chartPaths([]string{"start"}, cavern)
	fmt.Printf("%d\n", len(paths))
	os.Exit(0)
}

func isSmall(name string) bool {
	for _, c := range name {
		if c > 96 && c < 123 {
			continue
		}
		return false
	}
	return true
}

func chartPaths(trail []string, cavern map[string]node) [][]string {
	var chart [][]string
	if len(trail) == 0 {
		return nil
	}
	c := trail[len(trail)-1]
	if c == "end" {
		return append(chart, trail)
	}
	cur := cavern[c]
CheckNeighbors:
	for _, n := range cur.neighbors {
		if n == c {
			continue
		}
		next := cavern[n]
		visits := visited(n, trail)
		if len(visits) > 0 && next.small {
			continue
		}
		for _, v := range visits {
			if v == 0 {
				continue CheckNeighbors
			}
			if trail[v-1] == c {
				continue CheckNeighbors
			}
		}
		chart = append(chart, chartPaths(append(trail, n), cavern)...)
	}
	return chart
}

func visited(name string, trail []string) []int {
	var visits []int
	for i, cave := range trail {
		if cave == name {
			visits = append(visits, i)
		}
	}
	return visits
}
