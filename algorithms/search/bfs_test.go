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

func TestBFS2(t *testing.T) {
	TestBFS(t)
}
