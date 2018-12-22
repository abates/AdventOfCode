package bfs

import (
	"github.com/abates/AdventOfCode/graph"
)

type PathNode struct {
	path []graph.Node
}

func (p *PathNode) Node() graph.Node {
	return p.path[len(p.path)-1]
}

type NodeQueue struct {
	queue []*PathNode
}

func (n *NodeQueue) Len() int {
	return len(n.queue)
}

func (n *NodeQueue) Nodes() []graph.Node {
	nodes := make([]graph.Node, len(n.queue))
	for i, item := range n.queue {
		nodes[i] = item.Node()
	}
	return nodes
}

func (n *NodeQueue) Peek() (node *PathNode) {
	if len(n.queue) > 0 {
		node = n.queue[0]
	}
	return
}

func (n *NodeQueue) Push(nodes ...*PathNode) {
	n.queue = append(n.queue, nodes...)
}

func (n *NodeQueue) Shift() (node *PathNode) {
	if len(n.queue) > 0 {
		node = n.queue[0]
		n.queue = n.queue[1:]
	}
	return
}

type VisitFn func(int, []graph.Node) bool

func TraverseLevel(rootNode graph.Node, level int) []graph.Node {
	nodes := make([]graph.Node, 0)
	Traverse(rootNode, func(l int, path []graph.Node) bool {
		if l == level {
			nodes = append(nodes, path[len(path)-1])
		}
		return l == level+1
	})

	return nodes
}

func Height(rootNode graph.Node) int {
	level := 0
	Traverse(rootNode, func(l int, node []graph.Node) bool {
		level = l
		return false
	})
	return level
}

func Find(rootNode graph.Node, node graph.Node) (path []graph.Node) {
	Traverse(rootNode, func(l int, p []graph.Node) bool {
		if idable1, ok := p[len(p)-1].(graph.IDAble); ok {
			if idable2, ok := node.(graph.IDAble); ok {
				if idable1.ID() == idable2.ID() {
					path = p
					return true
				}
			}
		}

		if p[len(p)-1] == node {
			path = p
			return true
		}
		return false
	})
	return path
}

func Traverse(rootNode graph.Node, visit VisitFn) {
	visited := make(map[interface{}]struct{})
	q := NodeQueue{
		queue: make([]*PathNode, 0),
	}

	position := rootNode
	q.Push(&PathNode{
		path: []graph.Node{position},
	})
	if idable, ok := position.(graph.IDAble); ok {
		visited[idable.ID()] = struct{}{}
	} else {
		visited[position] = struct{}{}
	}

	for q.Len() > 0 {
		level := len(q.Peek().path) - 1
		if visit(level, q.Peek().path) {
			return
		}

		nextNode := q.Shift()
		position = nextNode.Node()

		for _, edge := range position.Edges() {
			node := edge.Neighbor()
			var id interface{}
			id = node
			if idable, ok := node.(graph.IDAble); ok {
				id = idable.ID()
			}
			if _, found := visited[id]; !found {
				n := &PathNode{
					path: make([]graph.Node, len(nextNode.path)+1),
				}
				copy(n.path, nextNode.path)
				n.path[len(n.path)-1] = node
				q.Push(n)
				visited[id] = struct{}{}
			}
		}
	}
}
