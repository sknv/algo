package binary_tree

type BinaryTreeNode struct {
	Value interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func (t *BinaryTree) Invert() {
	invert(t.Root)
}

func (t *BinaryTree) TraversePreOrder(visit func(value interface{})) {
	traversePreOrder(t.Root, visit)
}

func (t *BinaryTree) TraverseInOrder(visit func(value interface{})) {
	traverseInOrder(t.Root, visit)
}

func (t *BinaryTree) TraversePostOrder(visit func(value interface{})) {
	traversePostOrder(t.Root, visit)
}

func invert(node *BinaryTreeNode) {
	if node == nil {
		return
	}

	node.Left, node.Right = node.Right, node.Left
	invert(node.Left)
	invert(node.Right)
}

func traversePreOrder(node *BinaryTreeNode, visit func(value interface{})) {
	if node == nil {
		return
	}

	visit(node.Value)
	traversePreOrder(node.Left, visit)
	traversePreOrder(node.Right, visit)
}

func traverseInOrder(node *BinaryTreeNode, visit func(value interface{})) {
	if node == nil {
		return
	}

	traverseInOrder(node.Left, visit)
	visit(node.Value)
	traverseInOrder(node.Right, visit)
}

func traversePostOrder(node *BinaryTreeNode, visit func(value interface{})) {
	if node == nil {
		return
	}

	traversePostOrder(node.Left, visit)
	traversePostOrder(node.Right, visit)
	visit(node.Value)
}
