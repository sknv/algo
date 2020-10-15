package search

func DFS(root *GraphNode, visit func(value *GraphNode)) {
	dfs(root, visit, make(map[*GraphNode]struct{}))
}

func dfs(root *GraphNode, visit func(value *GraphNode), visited map[*GraphNode]struct{}) {
	if root == nil {
		return
	}

	visit(root)
	visited[root] = struct{}{}

	for _, n := range root.Nodes {
		if _, found := visited[n]; !found {
			dfs(n, visit, visited)
		}
	}
}
