package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/abates/AdventOfCode/2017/graph"
)

func buildGraph(input string) *graph.Graph {
	graph := graph.New()

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		fields := strings.Split(line, "<->")
		pid := strings.TrimSpace(fields[0])

		program := graph.FindOrCreateVertex(pid)

		for _, field := range strings.Split(fields[1], ",") {
			field = strings.TrimSpace(field)
			vertex := graph.FindOrCreateVertex(field)
			program.Connect(vertex)
		}
	}

	return graph
}

func main() {
	f, _ := os.Open("../input.txt")
	b, _ := ioutil.ReadAll(f)
	graph := buildGraph(string(b))
	program := graph.FindOrCreateVertex("0")
	if program != nil {
		connected := program.Connected()
		fmt.Printf("Vertex 0 has %d directly connected programs\n", len(program.Edges))
		fmt.Printf("Num: %d\n", len(connected))
	} else {
		fmt.Printf("Something went wrong and we got nil for program zero\n")
	}

	numGroups := 0
	assignedVertices := make(map[string]bool)
	for _, vertex := range graph.Vertices {
		if found := assignedVertices[vertex.ID]; !found {
			numGroups++
			for _, member := range vertex.Connected() {
				assignedVertices[member.ID] = true
			}
			assignedVertices[vertex.ID] = true
		}
	}
	fmt.Printf("Total number of groups %d\n", numGroups)
}
