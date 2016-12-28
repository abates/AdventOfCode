package alg

type Edge interface {
	Weight() int
	Neighbor() string
}

type BasicEdge struct {
	weight   int
	neighbor string
}

func NewBasicEdge(weight int, neighbor string) *BasicEdge {
	return &BasicEdge{weight, neighbor}
}

func (b *BasicEdge) Neighbor() string {
	return b.neighbor
}

func (b *BasicEdge) Weight() int {
	return b.weight
}

type GraphNode interface {
	ID() string
	Edges() []Edge
}

type BasicGraphNode struct {
	id    string
	edges []Edge
}

func NewBasicGraphNode(id string) *BasicGraphNode {
	return &BasicGraphNode{id, nil}
}

func (b *BasicGraphNode) ID() string {
	return b.id
}

func (b *BasicGraphNode) AddEdge(edge Edge) {
	b.edges = append(b.edges, edge)
}

func (b *BasicGraphNode) Edges() []Edge {
	return b.edges
}

type Graph interface {
	Nodes() []GraphNode
}

type BasicGraph struct {
	nodes map[string]GraphNode
}

func NewBasicGraph() *BasicGraph {
	return &BasicGraph{make(map[string]GraphNode)}
}

func (b *BasicGraph) AddNode(node GraphNode) {
	b.nodes[node.ID()] = node
}

func (b *BasicGraph) GetNode(id string) GraphNode {
	return b.nodes[id]
}

func (b *BasicGraph) Nodes() []GraphNode {
	nodes := make([]GraphNode, len(b.nodes))
	i := 0
	for _, node := range b.nodes {
		nodes[i] = node
		i++
	}
	return nodes
}

func SPFAll(graph Graph) map[string]map[string]int {
	distances := make(map[string]map[string]int)
	for _, node := range graph.Nodes() {
		distances[node.ID()] = SPF(graph, node)
	}
	return distances
}

func SPF(graph Graph, rootNode GraphNode) map[string]int {
	q := make(map[string]GraphNode, 0)
	distances := make(map[string]int)

	for _, node := range graph.Nodes() {
		q[node.ID()] = node
	}

	distances[rootNode.ID()] = 0
	for len(q) > 0 {
		var next GraphNode
		for id, node := range q {
			if next == nil {
				next = node
			} else {
				if d1, found := distances[id]; found {
					if d2, found := distances[next.ID()]; found {
						if d1 < d2 {
							next = node
						}
					} else {
						next = node
					}
				}
			}
		}

		delete(q, next.ID())

		for _, edge := range next.Edges() {
			alt := distances[next.ID()] + edge.Weight()
			if d, found := distances[edge.Neighbor()]; !found || alt < d {
				distances[edge.Neighbor()] = alt
			}
		}
	}

	return distances
}
