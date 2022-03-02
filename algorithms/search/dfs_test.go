package search

import (
	"reflect"
	"testing"
)

func TestDFS(t *testing.T) {
	root := makeGraph()

	expected := []int{0, 1, 3, 2, 4, 5}
	var values []int
	DFS(root, func(value *GraphNode) {
		values = append(values, value.Value.(int))
	})

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("expect bfs values to be [%v], got instead [%v]", expected, values)
	}
}

func TestDFS2(t *testing.T) {
	TestDFS(t)
}
