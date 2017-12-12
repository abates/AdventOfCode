package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	Node1 *Node
	Node2 *Node
}

type Node struct {
	ID       int
	Vertices []*Vertex
}

func (n *Node) Connect(other *Node) {
	vertex := &Vertex{n, other}
	n.Vertices = append(n.Vertices, vertex)
	other.Vertices = append(other.Vertices, vertex)
}

func (n *Node) Connected() map[int]*Node {
	visited := make(map[int]*Node)
	visited[n.ID] = n
	queue := []*Node{n}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, vertex := range current.Vertices {
			for _, node := range []*Node{vertex.Node1, vertex.Node2} {
				if node != current {
					if _, found := visited[node.ID]; !found {
						queue = append(queue, node)
						visited[node.ID] = node
					}
				}
			}
		}
	}
	return visited
}

type Graph struct {
	Nodes map[int]*Node
}

func (g *Graph) Find(id int) *Node {
	if g.Nodes == nil {
		g.Nodes = make(map[int]*Node)
	}

	node, found := g.Nodes[id]
	if !found {
		node = &Node{ID: id}
		g.Nodes[id] = node
	}

	return node
}

func buildGraph(input string) *Graph {
	graph := &Graph{}

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		fields := strings.Split(line, "<->")
		pid, _ := strconv.Atoi(strings.TrimSpace(fields[0]))

		program := graph.Find(pid)

		for _, field := range strings.Split(fields[1], ",") {
			field = strings.TrimSpace(field)
			c, _ := strconv.Atoi(field)
			node := graph.Find(c)
			program.Connect(node)
		}
	}

	return graph
}

func main() {
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	graph := buildGraph(string(b))
	program := graph.Find(0)
	if program != nil {
		connected := program.Connected()
		fmt.Printf("Node 0 has %d directly connected programs\n", len(program.Vertices))
		fmt.Printf("Num: %d\n", len(connected))
	} else {
		fmt.Printf("Something went wrong and we got nil for program zero\n")
	}

	numGroups := 0
	assignedNodes := make(map[int]bool)
	for _, node := range graph.Nodes {
		if found := assignedNodes[node.ID]; !found {
			numGroups++
			for _, member := range node.Connected() {
				assignedNodes[member.ID] = true
			}
			assignedNodes[node.ID] = true
		}
	}
	fmt.Printf("Total number of groups %d\n", numGroups)
}
