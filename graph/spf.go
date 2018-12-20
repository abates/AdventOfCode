package graph

type Edge interface {
	Weight() int
	Neighbor() Node
}

type Node interface {
	Edges() []Edge
}

type Graph interface {
	Nodes() []Node
}

func SPFAll(graph Graph) map[Node]map[Node]int {
	distances := make(map[Node]map[Node]int)
	for _, node := range graph.Nodes() {
		distances[node] = SPF(graph, node)
	}
	return distances
}

func SPF(graph Graph, rootNode Node) map[Node]int {
	q := make(map[Node]Node, 0)
	distances := make(map[Node]int)

	for _, node := range graph.Nodes() {
		q[node] = node
	}

	distances[rootNode] = 0
	for len(q) > 0 {
		var next Node
		for id, node := range q {
			if next == nil {
				next = node
			} else {
				if d1, found := distances[id]; found {
					if d2, found := distances[next]; found {
						if d1 < d2 {
							next = node
						}
					} else {
						next = node
					}
				}
			}
		}

		delete(q, next)

		for _, edge := range next.Edges() {
			alt := distances[next] + edge.Weight()
			if d, found := distances[edge.Neighbor()]; !found || alt < d {
				distances[edge.Neighbor()] = alt
			}
		}
	}

	return distances
}
