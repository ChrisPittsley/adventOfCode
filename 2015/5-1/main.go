package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
)

type word struct {
	vowels int
	dubs, bad bool
}

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	var niceWords int
	var currentWord word
	for i, char := range data {
		if i == len(data) - 1 {
			switch char {
			case 101, 105, 111, 117:
				currentWord.vowels += 1
			}
			if currentWord.vowels > 2 && currentWord.dubs && !currentWord.bad {
				niceWords += 1
			}
			continue
		}
		switch char {
		case 97:
			currentWord.vowels += 1
			currentWord.bad = (data[i+1] == 98) || currentWord.bad
		case 99:
			currentWord.bad = (data[i+1] == 100) || currentWord.bad
		case 112:
			currentWord.bad = (data[i+1] == 113) || currentWord.bad
		case 120:
			currentWord.bad = (data[i+1] == 121) || currentWord.bad
		case 101, 105, 111, 117:
			currentWord.vowels += 1
		case 10:
			if currentWord.vowels > 2 && currentWord.dubs && !currentWord.bad {
				niceWords += 1
			}
			currentWord = word{}
			continue
		}
		currentWord.dubs = (char == data[i+1]) || currentWord.dubs
	}
	fmt.Printf("%d\n", niceWords)
	os.Exit(0)
}
