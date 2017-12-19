package graph

type Edge struct {
	ID      string
	Vertex1 *Vertex
	Vertex2 *Vertex
}

type Vertex struct {
	ID    string
	Edges []*Edge
}

func NewVertex(id string) *Vertex {
	return &Vertex{ID: id}
}

func (v *Vertex) Connect(other *Vertex) *Edge {
	edge := &Edge{Vertex1: v, Vertex2: other}
	v.Edges = append(v.Edges, edge)
	other.Edges = append(other.Edges, edge)
	return edge
}

func (v *Vertex) Connected() map[string]*Vertex {
	visited := make(map[string]*Vertex)
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
	Vertices map[string]*Vertex
}

func New() *Graph {
	return &Graph{
		Vertices: make(map[string]*Vertex),
	}
}

func (g *Graph) FindOrCreateVertex(id string) *Vertex {
	vertex, found := g.Vertices[id]
	if !found {
		vertex = &Vertex{ID: id}
		g.Vertices[id] = vertex
	}

	return vertex
}
