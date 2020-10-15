package search

type GraphNode struct {
	Value interface{}
	Nodes []*GraphNode
}
