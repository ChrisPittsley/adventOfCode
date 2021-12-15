package main

import (
	"adventOfCode/lib"
	"fmt"
	"os"
)

type coordinates struct {
	x, y int
}

type node struct {
	cost, totalCost int
}

func main() {
	data, err := lib.GetInput()
	if err != nil {
		lib.ErrorOut(err)
	}
	var cavern = make(map[coordinates]node)
	var unvisited = make(map[coordinates]struct{})
	var pos coordinates
	for i, b := range data {
		switch b {
		case 48:
			cavern[pos] = node{cost: 0, totalCost: -1}
		case 49:
			cavern[pos] = node{cost: 1, totalCost: -1}
		case 50:
			cavern[pos] = node{cost: 2, totalCost: -1}
		case 51:
			cavern[pos] = node{cost: 3, totalCost: -1}
		case 52:
			cavern[pos] = node{cost: 4, totalCost: -1}
		case 53:
			cavern[pos] = node{cost: 5, totalCost: -1}
		case 54:
			cavern[pos] = node{cost: 6, totalCost: -1}
		case 55:
			cavern[pos] = node{cost: 7, totalCost: -1}
		case 56:
			cavern[pos] = node{cost: 8, totalCost: -1}
		case 57:
			cavern[pos] = node{cost: 9, totalCost: -1}
		case 10:
			pos.y += 1
			pos.x = 0
			continue
		default:
			lib.ErrorOut(fmt.Errorf("bad input at position %d", i))
		}
		unvisited[pos] = struct{}{}
		pos.x += 1
	}
	var cur, end coordinates
	for p := range cavern {
		if p.x >= end.x && p.y >= end.y {
			end = p
		}
	}
	var queue = make(map[coordinates]struct{})
	start := cavern[cur]
	start.totalCost = 0
	cavern[cur] = start
	for len(unvisited) > 0 {
		neighbors := []coordinates{
			{cur.x, cur.y + 1},
			{cur.x, cur.y - 1},
			{cur.x + 1, cur.y},
			{cur.x - 1, cur.y},
		}
		for _, n := range neighbors {
			if _, ok := unvisited[n]; !ok {
				continue
			}
			nbr, ok := cavern[n]
			if !ok {
				continue
			} else if nbr.totalCost == -1 || cavern[cur].totalCost+nbr.cost < nbr.totalCost {
				nbr.totalCost = cavern[cur].totalCost + nbr.cost
			}
			cavern[n] = nbr
			queue[n] = struct{}{}
		}
		delete(queue, cur)
		delete(unvisited, cur)
		min := -1
		for next := range queue {
			if n, ok := cavern[next]; !ok || n.totalCost == -1 {
				continue
			} else if n.totalCost < min || min == -1 {
				min = n.totalCost
				cur = next
			}
		}
	}
	fmt.Printf("%d\n", cavern[end].totalCost)
	os.Exit(0)
}
