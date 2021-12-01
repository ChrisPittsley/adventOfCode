package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
)

type word struct {
	twoPair, dubs bool
}

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	var niceWords int
	var currentWord word
	pairs := make(map[string]int)
	for i, char := range data {
		if char == 10 || i == len(data) - 1 {
			if currentWord.twoPair && currentWord.dubs {
				niceWords += 1
			}
			currentWord = word{}
			pairs = make(map[string]int)
			continue
		}
		if i < len(data) - 2 && data[i+1] != 10 {
			currentWord.dubs = (char == data[i+2]) || currentWord.dubs
		}
		if pos, ok := pairs[string(data[i:i+2])]; !ok {
			pairs[string(data[i:i+2])] = i
		} else {
			currentWord.twoPair = (pos != i-1) || currentWord.twoPair
		}
	}
	fmt.Printf("%d\n", niceWords)
	os.Exit(0)
}
