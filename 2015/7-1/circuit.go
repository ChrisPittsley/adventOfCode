package main

import (
	"strconv"
)

type Wire struct {
	output []*Gate
	signal uint16
	active bool
}

type Gate struct {
	input  [2]*Wire
	output *Wire
	op     func(in [2]uint16) uint16
}

type Circuit struct {
	graph map[string]*Wire
}

func (c *Circuit) Insert(s statement) {
	gate := c.newGate(s)
	gate.update()
}

func (c *Circuit) newGate(s statement) Gate {
	var gate Gate
	switch s.operation {
	case NOT:
		gate.op = not
	case AND:
		gate.op = and
	case OR:
		gate.op = or
	case LSHIFT:
		gate.op = lshift
	case RSHIFT:
		gate.op = rshift
	case NOOP:
		gate.op = noop
	}
	for n, i := range s.inputs {
		var w Wire
		switch {
		case isNumber(i):
			signal, _ := strconv.ParseUint(i, 10, 16)
			w.signal = uint16(signal)
			w.active = true
			gate.input[n] = &w
		case isWireName(i):
			if wire, ok := c.graph[i]; ok {
				wire.output = append(wire.output, &gate)
				gate.input[n] = wire
			} else {
				w.output = []*Gate{&gate}
				gate.input[n] = &w
				c.graph[i] = &w
			}
		}
		if s.operation == NOOP || s.operation == NOT {
			var dummy = Wire{active: true}
			gate.input[1] = &dummy
			break
		}
	}
	w := new(Wire)
	var ok bool
	if _, ok = c.graph[s.output]; !ok {
		c.graph[s.output] = w
	} else {
		w = c.graph[s.output]
	}
	gate.output = w
	return gate
}

func (g *Gate) update() {
	if !g.input[0].active || !g.input[1].active {
		return
	}
	g.output.signal = g.op([2]uint16{g.input[0].signal, g.input[1].signal})
	g.output.active = true
	for x := range g.output.output {
		g.output.output[x].update()
	}
	return
}

func and(in [2]uint16) uint16 {
	return in[0] & in[1]
}

func or(in [2]uint16) uint16 {
	return in[0] | in[1]
}

func lshift(in [2]uint16) uint16 {
	return in[0] << in[1]
}

func rshift(in [2]uint16) uint16 {
	return in[0] >> in[1]
}

func not(in [2]uint16) uint16 {
	return in[0] ^ 0xFFFF
}

func noop(in [2]uint16) uint16 {
	return in[0]
}
