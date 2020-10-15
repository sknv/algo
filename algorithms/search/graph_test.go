package search

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
