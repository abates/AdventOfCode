package bfs

import (
	"testing"

	"github.com/abates/AdventOfCode/coordinate"
	"github.com/abates/AdventOfCode/graph"
)

type Coordinate struct {
	*coordinate.Coordinate
}

func NewCoordinate(x, y int) *Coordinate {
	return &Coordinate{coordinate.New(x, y)}
}

func (c *Coordinate) ID() string {
	return c.String()
}

func (c *Coordinate) Edges() []graph.Edge {
	edges := make([]graph.Edge, 0)
	directions := []*coordinate.Coordinate{coordinate.New(0, -1), coordinate.New(1, 0), coordinate.New(0, 1), coordinate.New(-1, 0)}
	for _, direction := range directions {
		edges = append(edges, graph.NewBasicEdge(1, &Coordinate{c.Add(direction)}))
	}

	return edges
}

func (c *Coordinate) Equal(other *Coordinate) bool { return c.Coordinate.Equal(other.Coordinate) }
func TestTraverse(t *testing.T) {
	tests := []struct {
		startX      int
		startY      int
		destination *Coordinate
		result      int
	}{
		{1, 1, NewCoordinate(4, 4), 7},
	}

	for i, test := range tests {
		var result []graph.Node
		Traverse(NewCoordinate(test.startX, test.startY), func(l int, path []graph.Node) bool {
			result = path
			return result[len(result)-1].(*Coordinate).Equal(test.destination)
		})

		if len(result) != test.result {
			t.Errorf("Test %d: Unexpected path length.  Expected %d got %d", i, test.result, len(result))
		}
	}
}

func TestFindAt(t *testing.T) {
	tests := []struct {
		startX int
		startY int
		level  int
		result int
	}{
		{1, 1, 1, 4},
		{1, 1, 2, 8},
		{1, 1, 3, 12},
		{1, 1, 4, 16},
	}

	for i, test := range tests {
		result := TraverseLevel(NewCoordinate(1, 1), test.level)
		if len(result) != test.result {
			t.Errorf("Test %d: Unexpected path length.  Expected %d got %d %v", i, test.result, len(result), result)
		}
	}
}
