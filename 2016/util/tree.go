package util

type Node interface {
	ID() string
	Neighbors() []Node
	Equal(Node) bool
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

type Tree struct {
	position Node
	visited  map[string]struct{}
}

func NewTree(startNode Node) *Tree {
	return &Tree{
		position: startNode,
	}
}

func (t *Tree) VisitedNodes() []string {
	ids := make([]string, len(t.visited))
	i := 0
	for id, _ := range t.visited {
		ids[i] = id
		i++
	}

	return ids
}

func (t *Tree) Find(destination Node) []Node {
	_, path := t.find(destination, nil)
	return path
}

type StopFn func(int, Node) bool

func (t *Tree) FindAt(level int) []Node {
	q, _ := t.find(nil, func(l int, n Node) bool {
		return l == level+1
	})

	return q
}

func (t *Tree) FindUntil(cb StopFn) []Node {
	_, path := t.find(nil, cb)
	return path
}

func (t *Tree) Height(destination Node) int {
	level := 0
	t.find(nil, func(l int, node Node) bool {
		if destination.Equal(node) {
			level = l
		}
		return false
	})
	return level
}

func (t *Tree) find(destination Node, stop StopFn) ([]Node, []Node) {
	t.visited = make(map[string]struct{})
	q := NodeQueue{
		queue: make([]*PathNode, 0),
	}

	q.Push(&PathNode{
		path: []Node{t.position},
	})
	t.visited[t.position.ID()] = struct{}{}

	for q.Len() > 0 {
		level := len(q.Peek().path)
		if stop != nil && stop(level, t.position) {
			return q.Nodes(), q.Peek().path
		}

		nextNode := q.Shift()
		t.position = nextNode.Node()

		if t.position.Equal(destination) {
			return q.Nodes(), nextNode.path
		}

		for _, node := range t.position.Neighbors() {
			if _, found := t.visited[node.ID()]; !found {
				n := &PathNode{
					path: make([]Node, len(nextNode.path)+1),
				}
				copy(n.path, nextNode.path)
				n.path[len(n.path)-1] = node
				q.Push(n)
				t.visited[node.ID()] = struct{}{}
			}
		}
	}

	return q.Nodes(), nil
}
