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

type IDAble interface {
	ID() string
}
