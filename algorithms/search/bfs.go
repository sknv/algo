package search

import (
	"algo/data-structures/queue"
)

func BFS(root *GraphNode, visit func(value *GraphNode)) {
	if root == nil {
		return
	}

	q := queue.New()
	q.Enqueue(root)
	visited := map[*GraphNode]struct{}{
		root: {},
	}

	for !q.IsEmpty() {
		node := q.Dequeue().(*GraphNode)
		visit(node)
		for _, n := range node.Nodes {
			if _, found := visited[n]; !found {
				q.Enqueue(n)
				visited[n] = struct{}{}
			}
		}
	}
}
