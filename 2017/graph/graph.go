package graph

type Edge struct {
	Vertex1 *Vertex
	Vertex2 *Vertex
}

type Vertex struct {
	ID    int
	Edges []*Edge
}

func (v *Vertex) Connect(other *Vertex) {
	edge := &Edge{v, other}
	v.Edges = append(v.Edges, edge)
	other.Edges = append(other.Edges, edge)
}

func (v *Vertex) Connected() map[int]*Vertex {
	visited := make(map[int]*Vertex)
	visited[v.ID] = v
	queue := []*Vertex{v}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, edge := range current.Edges {
			for _, vertex := range []*Vertex{edge.Vertex1, edge.Vertex2} {
				if vertex != current {
					if _, found := visited[vertex.ID]; !found {
						queue = append(queue, vertex)
						visited[vertex.ID] = vertex
					}
				}
			}
		}
	}
	return visited
}

type Graph struct {
	Vertices map[int]*Vertex
}

func New() *Graph {
	return &Graph{
		Vertices: make(map[int]*Vertex),
	}
}

func (g *Graph) FindOrCreateVertex(id int) *Vertex {
	vertex, found := g.Vertices[id]
	if !found {
		vertex = &Vertex{ID: id}
		g.Vertices[id] = vertex
	}

	return vertex
}
