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
	var rules = make(map[string]rune)
	var polymer string
	for l, line := range strings.Split(string(data), "\n") {
		if l == 0 {
			polymer = line
			continue
		}
		if line == "" {
			continue
		}
		rule := strings.Split(line, " -> ")
		if len(rule) != 2 {
			lib.ErrorOut(fmt.Errorf("bad syntax on line %d", l))
		}
		if len(rule[0]) != 2 || len(rule[1]) != 1 {
			lib.ErrorOut(fmt.Errorf("bad syntax on line %d", l))
		}
		//rules[rule[0]] = string([]uint8{rule[0][0], rule[1][0], rule[0][1]})
		rules[rule[0]] = rune(rule[1][0])
	}
	for n := 0; n < 10; n += 1 {
		polymer = insert(polymer, rules)
	}
	var max, min int
	var buckets = make(map[rune]int)
	for _, char := range polymer {
		x := buckets[char]
		x += 1
		if x > max {
			max = x
		}
		buckets[char] = x
	}
	for _, x := range buckets {
		if x < min || min == 0 {
			min = x
		}
	}
	fmt.Printf("%d\n", max-min)
	os.Exit(0)
}

func insert(s string, rules map[string]rune) string {
	end := len(s) - 1
	var out []rune
	for i, c := range s {
		out = append(out, c)
		if i == end {
			break
		}
		pair := string([]rune{c, rune(s[i+1])})
		if ins, ok := rules[pair]; ok {
			out = append(out, ins)
		}
	}
	return string(out)
}
