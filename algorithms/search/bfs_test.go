package search

import (
	"reflect"
	"testing"
)

func TestBFS(t *testing.T) {
	root := makeGraph()

	expected := []int{0, 1, 4, 5, 3, 2}
	var values []int
	BFS(root, func(value *GraphNode) {
		values = append(values, value.Value.(int))
	})

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("expect bfs values to be [%v], got instead [%v]", expected, values)
	}
}

func makeGraph() *GraphNode {
	zero := GraphNode{Value: 0}
	one := GraphNode{Value: 1}
	two := GraphNode{Value: 2}
	three := GraphNode{Value: 3}
	four := GraphNode{Value: 4}
	five := GraphNode{Value: 5}

	zero.Nodes = []*GraphNode{
		&one, &four, &five,
	}
	one.Nodes = []*GraphNode{
		&three, &four,
	}
	two.Nodes = []*GraphNode{
		&one,
	}
	three.Nodes = []*GraphNode{
		&two, &four,
	}
	return &zero
}
