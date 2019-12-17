package main

import (
	"fmt"
	"strings"
)

func init() {
	d6 := &D6{}
	challenges[6] = &challenge{"Day 06", "input/day06.txt", d6}
}

type orbitNode struct {
	name   string
	center *orbitNode
	orbits []*orbitNode
}

func (on *orbitNode) IOCount(depth int) int {
	count := depth
	for _, o := range on.orbits {
		count += o.IOCount(depth + 1)
	}
	return count
}

func (on *orbitNode) PathToRoot() []*orbitNode {
	nodes := []*orbitNode{}
	for n := on.center; n != nil; n = n.center {
		nodes = append(nodes, n)
	}
	return nodes
}

func (on *orbitNode) OTCount(other *orbitNode) int {
	mp := on.PathToRoot()
	op := other.PathToRoot()
	for i, m := range mp {
		for j, o := range op {
			if m == o {
				return i + j
			}
		}
	}
	return -1
}

type D6 struct {
	nodeIndex map[string]*orbitNode
	orbitMap  *orbitNode
}

func (d6 *D6) parse(line string) error {
	if d6.orbitMap == nil {
		d6.nodeIndex = make(map[string]*orbitNode)
		d6.orbitMap = &orbitNode{name: "COM"}
		d6.nodeIndex["COM"] = d6.orbitMap
	}

	names := strings.Split(line, ")")

	center, found := d6.nodeIndex[names[0]]
	if !found {
		center = &orbitNode{name: names[0]}
		d6.nodeIndex[names[0]] = center
	}

	on, found := d6.nodeIndex[names[1]]
	if !found {
		on = &orbitNode{name: names[1]}
		d6.nodeIndex[names[1]] = on
	}
	on.center = center
	center.orbits = append(center.orbits, on)
	return nil
}

func (d6 *D6) part1() (string, error) {
	count := d6.orbitMap.IOCount(0)
	return fmt.Sprintf("IO Count: %d", count), nil
}

func (d6 *D6) part2() (string, error) {
	count := d6.nodeIndex["YOU"].OTCount(d6.nodeIndex["SAN"])
	return fmt.Sprintf("%d orbital transfers required", count), nil
}
