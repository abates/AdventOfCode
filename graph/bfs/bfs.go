package bfs

import (
	"github.com/abates/AdventOfCode/graph"
)

type Path struct {
	level int
	prev  *Path
	node  graph.Node
}

func (p *Path) Append(node graph.Node) *Path {
	return &Path{
		level: p.level + 1,
		prev:  p,
		node:  node,
	}
}

func (p *Path) Path() []graph.Node {
	path := []graph.Node{p.node}
	for p.prev != nil {
		path = append([]graph.Node{p.node}, path...)
		p = p.prev
	}
	return path
}

type NodeQueue struct {
	queue []*Path
}

func (n *NodeQueue) Len() int {
	return len(n.queue)
}

func (n *NodeQueue) Nodes() []graph.Node {
	nodes := make([]graph.Node, len(n.queue))
	for i, item := range n.queue {
		nodes[i] = item.node
	}
	return nodes
}

func (n *NodeQueue) Peek() (node *Path) {
	if len(n.queue) > 0 {
		node = n.queue[0]
	}
	return
}

func (n *NodeQueue) Push(nodes ...*Path) {
	n.queue = append(n.queue, nodes...)
}

func (n *NodeQueue) Shift() (node *Path) {
	if len(n.queue) > 0 {
		node = n.queue[0]
		n.queue = n.queue[1:]
	}
	return
}

type VisitFn func(int, *Path) bool

func TraverseLevel(rootNode graph.Node, level int) []graph.Node {
	nodes := make([]graph.Node, 0)
	Traverse(rootNode, func(l int, path *Path) bool {
		if l == level {
			nodes = append(nodes, path.node)
		}
		return l == level+1
	})

	return nodes
}

func Height(rootNode graph.Node) int {
	level := 0
	Traverse(rootNode, func(l int, path *Path) bool {
		level = l
		return false
	})
	return level
}

func Find(rootNode graph.Node, node graph.Node) []graph.Node {
	var path *Path
	Traverse(rootNode, func(l int, p *Path) bool {
		if idable1, ok := p.node.(graph.IDAble); ok {
			if idable2, ok := node.(graph.IDAble); ok {
				if idable1.ID() == idable2.ID() {
					path = p
					return true
				}
			}
		}

		if p.node == node {
			path = p
			return true
		}
		return false
	})
	return path.Path()
}

func Traverse(rootNode graph.Node, visit VisitFn) {
	visited := make(map[interface{}]struct{})
	q := NodeQueue{
		queue: make([]*Path, 0),
	}

	position := rootNode
	q.Push(&Path{node: position, level: 1})
	if idable, ok := position.(graph.IDAble); ok {
		visited[idable.ID()] = struct{}{}
	} else {
		visited[position] = struct{}{}
	}

	for q.Len() > 0 {
		level := q.Peek().level - 1
		if visit(level, q.Peek()) {
			return
		}

		nextNode := q.Shift()
		position = nextNode.node

		for _, edge := range position.Edges() {
			node := edge.Neighbor()
			var id interface{}
			id = node
			if idable, ok := node.(graph.IDAble); ok {
				id = idable.ID()
			}
			if _, found := visited[id]; !found {
				p := nextNode.Append(node)
				q.Push(p)
				visited[id] = struct{}{}
			}
		}
	}
}
