package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Node struct {
	metadata []int
	children []*Node
}

func (n *Node) Parse(values []int) []int {
	numChildren := values[0]
	numMetadata := values[1]
	values = values[2:]
	for ; numChildren > 0; numChildren-- {
		child := &Node{}
		values = child.Parse(values)
		n.children = append(n.children, child)
	}

	for ; numMetadata > 0; numMetadata-- {
		n.metadata = append(n.metadata, values[0])
		values = values[1:]
	}
	return values
}

func (n *Node) MetadataSum1() int {
	sum := 0
	for _, m := range n.metadata {
		sum += m
	}
	for _, child := range n.children {
		sum += child.MetadataSum1()
	}
	return sum
}

func (n *Node) MetadataSum2() int {
	sum := 0
	if len(n.children) == 0 {
		for _, m := range n.metadata {
			sum += m
		}
	} else {
		for _, m := range n.metadata {
			if m <= len(n.children) {
				sum += n.children[m-1].MetadataSum2()
			}
		}
	}
	return sum
}

func main() {
	input, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic("Error: " + err.Error())
	}

	root := &Node{}
	values := []int{}
	for _, field := range strings.Fields(string(input)) {
		n, err := strconv.Atoi(field)
		if err != nil {
			panic("Error: " + err.Error())
		}
		values = append(values, n)
	}
	root.Parse(values)

	fmt.Printf("Metadata Sum 1: %d\n", root.MetadataSum1())
	fmt.Printf("Metadata Sum 2: %d\n", root.MetadataSum2())
}
