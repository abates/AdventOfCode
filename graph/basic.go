package graph

type BasicEdge struct {
	weight   int
	neighbor Node
}

func (b *BasicEdge) Neighbor() Node {
	return b.neighbor
}

func (b *BasicEdge) Weight() int {
	return b.weight
}

type BasicNode struct {
	id    string
	edges []Edge
}

func (b *BasicNode) AddEdge(edge Edge) {
	b.edges = append(b.edges, edge)
}

func (b *BasicNode) Edges() []Edge {
	return b.edges
}

type BasicGraph struct {
	nodes map[string]*BasicNode
}

func NewBasicGraph() *BasicGraph {
	return &BasicGraph{make(map[string]*BasicNode)}
}

func (b *BasicGraph) GetNode(id string) *BasicNode {
	return b.nodes[id]
}

func (b *BasicGraph) AddDirectedEdge(id1, id2 string, weight int) {
	node1, found := b.nodes[id1]
	if !found {
		node1 = &BasicNode{id: id1}
		if b.nodes == nil {
			b.nodes = make(map[string]*BasicNode)
		}
		b.nodes[id1] = node1
	}
	node2, found := b.nodes[id2]
	if !found {
		node2 = &BasicNode{id: id2}
		b.nodes[id2] = node2
	}
	node1.AddEdge(&BasicEdge{weight, node2})
}

func (b *BasicGraph) Nodes() []Node {
	nodes := make([]Node, len(b.nodes))
	i := 0
	for _, node := range b.nodes {
		nodes[i] = node
		i++
	}
	return nodes
}
