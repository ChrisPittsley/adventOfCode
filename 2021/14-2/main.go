package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	var rules = make(map[string][]string)
	var polymer string
	for l, line := range strings.Split(string(data), "\n") {
		if l == 0 {
			polymer = line
			continue
		}
		if line == "" {
			continue
		}
		stmt := strings.Split(line, " -> ")
		if len(stmt) != 2 {
			lib.ErrorOut(fmt.Errorf("bad syntax on line %d", l))
		}
		if len(stmt[0]) != 2 || len(stmt[1]) != 1 {
			lib.ErrorOut(fmt.Errorf("bad syntax on line %d", l))
		}
		rules[stmt[0]] = []string{string([]uint8{stmt[0][0], stmt[1][0]}), string([]uint8{stmt[1][0], stmt[0][1]})}
	}
	var pairs = make(map[string]uint)
	for i := 0; i < len(polymer)-1; i += 1 {
		pairs[polymer[i:i+2]] += 1
	}
	for i := 0; i < 40; i += 1 {
		newPairs := make(map[string]uint)
		for pair, count := range pairs {
			if rule, ok := rules[pair]; ok {
				for _, p := range rule {
					newPairs[p] += count
				}
				delete(pairs, pair)
			}
		}
		for pair, count := range newPairs {
			pairs[pair] += count
		}
	}
	var letters = make(map[uint8]uint)
	for pair, count := range pairs {
		letters[pair[0]] += count
	}
	letters[polymer[len(polymer)-1]] += 1
	var max, min uint
	for _, count := range letters {
		if count > max {
			max = count
		}
		if count < min || min == 0 {
			min = count
		}
	}
	fmt.Printf("%d\n", max-min)
	os.Exit(0)
}
