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
	var circuit Circuit
	circuit.graph = make(map[string]*Wire)
	for l, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}
		stmt, err := parse(line)
		if err != nil {
			lib.ErrorOut(fmt.Errorf("error on line %d '%s': %v", l, line, err))
		}
		circuit.Insert(stmt)
	}
	if a, ok := circuit.graph["a"]; !ok {
		lib.ErrorOut(fmt.Errorf("error: no value assigned to output 'a'"))
	} else {
		fmt.Printf("%d\n", a.signal)
	}
	os.Exit(0)
}
