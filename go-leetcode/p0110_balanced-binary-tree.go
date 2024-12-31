package leetcode

// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

func isBalanced(root *IntTreeNode) bool {
	var dfs func(node *IntTreeNode) (bool, int)
	dfs = func(node *IntTreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}

		leftBalanced, leftHeight := dfs(node.Left)
		rightBalanced, rightHeight := dfs(node.Right)

		balanced := abs(leftHeight-rightHeight) <= 1 &&
			leftBalanced &&
			rightBalanced

		return balanced, max(leftHeight, rightHeight) + 1
	}

	result, _ := dfs(root)

	return result
}
