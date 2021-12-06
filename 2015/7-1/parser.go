package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	NOT    = "NOT "
	AND    = " AND "
	OR     = " OR "
	LSHIFT = " LSHIFT "
	RSHIFT = " RSHIFT "
	NOOP   = "NOOP"
)

type statement struct {
	output, operation string
	inputs            [2]string
}

func parse(in string) (statement, error) {
	var stmt statement
	sides := strings.Split(in, " -> ")
	if len(sides) != 2 {
		return statement{}, fmt.Errorf("bad syntax: %s", in)
	}
	if !isWireName(sides[1]) {
		return statement{}, fmt.Errorf("bad output name: %s", sides[1])
	}
	stmt.output = sides[1]
	switch {
	case strings.HasPrefix(sides[0], NOT):
		s := strings.TrimPrefix(sides[0], NOT)
		_, err := strconv.Atoi(s)
		if err != nil && !isWireName(s) {
			return statement{}, fmt.Errorf("bad output name: %s", s)
		}
		stmt.operation, stmt.inputs[0] = NOT, s
	case len(strings.Split(sides[0], AND)) == 2:
		oprnds := strings.Split(sides[0], AND)
		for x, s := range oprnds {
			_, err := strconv.Atoi(s)
			if err != nil && !isWireName(s) {
				return statement{}, fmt.Errorf("bad output name: %s", s)
			}
			stmt.inputs[x] = s
		}
		stmt.operation = AND
	case len(strings.Split(sides[0], OR)) == 2:
		oprnds := strings.Split(sides[0], OR)
		for x, s := range oprnds {
			_, err := strconv.Atoi(s)
			if err != nil && !isWireName(s) {
				return statement{}, fmt.Errorf("bad output name: %s", s)
			}
			stmt.inputs[x] = s
		}
		stmt.operation = OR
	case len(strings.Split(sides[0], LSHIFT)) == 2:
		oprnds := strings.Split(sides[0], LSHIFT)
		for x, s := range oprnds {
			_, err := strconv.Atoi(s)
			if err != nil && !isWireName(s) {
				return statement{}, fmt.Errorf("bad output name: %s", s)
			}
			stmt.inputs[x] = s
		}
		stmt.operation = LSHIFT
	case len(strings.Split(sides[0], RSHIFT)) == 2:
		oprnds := strings.Split(sides[0], RSHIFT)
		for x, s := range oprnds {
			_, err := strconv.Atoi(s)
			if err != nil && !isWireName(s) {
				return statement{}, fmt.Errorf("bad output name: %s", s)
			}
			stmt.inputs[x] = s
		}
		stmt.operation = RSHIFT
	default:
		_, err := strconv.Atoi(sides[0])
		if err != nil && !isWireName(sides[0]) {
			return statement{}, fmt.Errorf("bad output name: %s", sides[0])
		}
		stmt.inputs[0] = sides[0]
		stmt.operation = NOOP
	}
	return stmt, nil
}

func isWireName(s string) bool {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	for _, r := range strings.ToLower(s) {
		if strings.ContainsRune(alphabet, r) {
			continue
		} else {
			return false
		}
	}
	return true
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
