package alg

type Node interface {
	ID() string
	Neighbors() []Node
}

type PathNode struct {
	path []Node
}

func (p *PathNode) Node() Node {
	return p.path[len(p.path)-1]
}

func (p *PathNode) String() string {
	return p.Node().ID()
}

type NodeQueue struct {
	queue []*PathNode
}

func (n *NodeQueue) Len() int {
	return len(n.queue)
}

func (n *NodeQueue) Nodes() []Node {
	nodes := make([]Node, len(n.queue))
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

type VisitFn func(int, []Node) bool

func TraverseLevel(rootNode Node, level int) []Node {
	nodes := make([]Node, 0)
	Traverse(rootNode, func(l int, path []Node) bool {
		if l == level {
			nodes = append(nodes, path[len(path)-1])
		}
		return l == level+1
	})

	return nodes
}

func Height(rootNode Node) int {
	level := 0
	Traverse(rootNode, func(l int, node []Node) bool {
		level = l
		return false
	})
	return level
}

func Find(rootNode Node, id string) (path []Node) {
	Traverse(rootNode, func(l int, p []Node) bool {
		if p[len(p)-1].ID() == id {
			path = p
			return true
		}
		return false
	})
	return path
}

func Traverse(rootNode Node, visit VisitFn) {
	visited := make(map[string]struct{})
	q := NodeQueue{
		queue: make([]*PathNode, 0),
	}

	position := rootNode
	q.Push(&PathNode{
		path: []Node{position},
	})
	visited[position.ID()] = struct{}{}

	for q.Len() > 0 {
		level := len(q.Peek().path) - 1
		if visit(level, q.Peek().path) {
			return
		}

		nextNode := q.Shift()
		position = nextNode.Node()

		for _, node := range position.Neighbors() {
			if _, found := visited[node.ID()]; !found {
				n := &PathNode{
					path: make([]Node, len(nextNode.path)+1),
				}
				copy(n.path, nextNode.path)
				n.path[len(n.path)-1] = node
				q.Push(n)
				visited[node.ID()] = struct{}{}
			}
		}
	}
}
