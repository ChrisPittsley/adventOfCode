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
	var current [9]int
	for _, n := range strings.Split(string(data), ",") {
		age, err := strconv.Atoi(strings.TrimSuffix(n, "\n"))
		if err != nil {
			lib.ErrorOut(err)
		}
		current[age] += 1
	}
	for day := 1; day <= 256; day += 1 {
		var spawn [9]int
		for age := 8; age > 0; age -= 1 {
			spawn[age-1] = current[age]
		}
		spawn[8] = current[0]
		spawn[6] += current[0]
		current = spawn
	}
	total := 0
	for _, x := range current {
		total += x
	}
	fmt.Printf("%d\n", total)
	os.Exit(0)
}
