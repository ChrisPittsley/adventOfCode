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
	var tile = make(map[coordinates]node)
	var unvisited = make(map[coordinates]struct{})
	var pos coordinates
	var rightEdge, bottomEdge int
	for i, b := range data {
		switch b {
		case 48:
			tile[pos] = node{cost: 0, totalCost: -1}
		case 49:
			tile[pos] = node{cost: 1, totalCost: -1}
		case 50:
			tile[pos] = node{cost: 2, totalCost: -1}
		case 51:
			tile[pos] = node{cost: 3, totalCost: -1}
		case 52:
			tile[pos] = node{cost: 4, totalCost: -1}
		case 53:
			tile[pos] = node{cost: 5, totalCost: -1}
		case 54:
			tile[pos] = node{cost: 6, totalCost: -1}
		case 55:
			tile[pos] = node{cost: 7, totalCost: -1}
		case 56:
			tile[pos] = node{cost: 8, totalCost: -1}
		case 57:
			tile[pos] = node{cost: 9, totalCost: -1}
		case 10:
			pos.y += 1
			pos.x = 0
			continue
		default:
			lib.ErrorOut(fmt.Errorf("bad input at position %d", i))
		}
		unvisited[pos] = struct{}{}
		if pos.x > rightEdge {
			rightEdge = pos.x
		}
		if pos.y > bottomEdge {
			bottomEdge = pos.y
		}
		pos.x += 1
	}
	var cavern = make(map[coordinates]node)
	var cur, end coordinates
	for p, n := range tile {
		for x := 0; x <= 4; x += 1 {
			newN := n
			newN.cost = n.cost + x
			if newN.cost > 9 {
				newN.cost -= 9
			}
			for y := 0; y <= 4; y += 1 {
				newP := coordinates{x: p.x + (rightEdge * x) + x, y: p.y + (bottomEdge * y) + y}
				cavern[newP] = newN
				unvisited[newP] = struct{}{}
				if newP.x >= end.x && newP.y >= end.y {
					end = newP
				}
				newN.cost += 1
				if newN.cost > 9 {
					newN.cost = 1
				}
			}
		}
	}
	start := cavern[cur]
	start.totalCost = 0
	cavern[cur] = start
	var queue = make(map[coordinates]struct{})
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
		delete(unvisited, cur)
		delete(queue, cur)
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
