package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
	"strings"
)

type digit struct {
	a, b, c, d, e, f, g bool
}

//descramble uses key to determine the mapping of each segment and returns the number represented by output.
//Frequency of each segment
//a: 8
//b: 6
//c: 8
//d: 7
//e: 4
//f: 9
//g: 7
//
//We know that 'a' and 'g' don't appear in the number 4, but 'c' and and 'd' do. We also know that 4 is the only number
//with 4 segments, therefore we can easily differentiate between 'a' and 'c' or 'd' and 'g'. Mapping the rest of the
//segments is easy because they all have unique frequencies.
func descramble(key, output []string) (int, error) {
	if len(key) != 10 {
		return 0, fmt.Errorf("not enough digits in key: 10 required, recieved %d", len(key))
	}
	var freq = map[rune]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0}
	var mapping = make(map[rune]rune)
	var four string
	for _, d := range key {
		if len(d) == 4 {
			four = d
		}
		for _, segment := range d {
			if _, ok := freq[segment]; !ok {
				return 0, fmt.Errorf("bad segment ID: %s", string(segment))
			}
			freq[segment] += 1
		}
	}
	var ac, dg []rune
	for segment, f := range freq {
		switch f {
		case 8:
			ac = append(ac, segment)
		case 7:
			dg = append(dg, segment)
		case 6:
			mapping[segment] = 'b'
		case 4:
			mapping[segment] = 'e'
		case 9:
			mapping[segment] = 'f'
		}
	}
	for _, segment := range ac {
		if strings.ContainsRune(four, segment) {
			mapping[segment] = 'c'
		} else {
			mapping[segment] = 'a'
		}
	}
	for _, segment := range dg {
		if strings.ContainsRune(four, segment) {
			mapping[segment] = 'd'
		} else {
			mapping[segment] = 'g'
		}
	}
	place := 1
	var out int
	for pos := len(output) - 1; pos >= 0; pos -= 1 {
		dgt, err := decode(mapping, output[pos])
		if err != nil {
			return 0, err
		}
		out += dgt * place
		place *= 10
	}
	return out, nil
}

//decode returns the integer represented by valid scrambled digit d, given a valid map
func decode(mapping map[rune]rune, in string) (int, error) {
	var numbers = [10]digit{
		{a: true, b: true, c: true, e: true, f: true, g: true},
		{c: true, f: true},
		{a: true, c: true, d: true, e: true, g: true},
		{a: true, c: true, d: true, f: true, g: true},
		{b: true, c: true, d: true, f: true},
		{a: true, b: true, d: true, f: true, g: true},
		{a: true, b: true, d: true, e: true, f: true, g: true},
		{a: true, c: true, f: true},
		{a: true, b: true, c: true, d: true, e: true, f: true, g: true},
		{a: true, b: true, c: true, d: true, f: true, g: true},
	}
	if len(mapping) != 7 {
		return 0, fmt.Errorf("not enough information for decoding")
	}
	var dgt digit
	for _, s := range in {
		switch mapping[s] {
		case 'a':
			dgt.a = true
		case 'b':
			dgt.b = true
		case 'c':
			dgt.c = true
		case 'd':
			dgt.d = true
		case 'e':
			dgt.e = true
		case 'f':
			dgt.f = true
		case 'g':
			dgt.g = true
		}
	}
	for n := range numbers {
		if numbers[n] == dgt {
			return n, nil
		}
	}
	return 0, fmt.Errorf("could not decode digit %s: %v", in, dgt)
}

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	sum := 0
	for l, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " | ")
		if len(parts) != 2 {
			lib.ErrorOut(fmt.Errorf("error on line %d: bad syntax: %s", l, line))
		}
		key := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")
		out, err := descramble(key, output)
		if err != nil {
			lib.ErrorOut(fmt.Errorf("error on line %d: %v", l, err))
		}
		sum += out
	}
	fmt.Printf("%d\n", sum)
	os.Exit(0)
}
