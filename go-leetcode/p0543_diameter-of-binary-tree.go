package leetcode

// https://leetcode.com/problems/diameter-of-binary-tree/description/

func diameterOfBinaryTree(root *IntTreeNode) int {
	var dfs func(node *IntTreeNode, result *int) int
	dfs = func(node *IntTreeNode, result *int) int {
		if node == nil {
			return 0
		}

		leftDepth := dfs(node.Left, result)
		rightDepth := dfs(node.Right, result)

		*result = max(*result, leftDepth+rightDepth)

		return max(leftDepth, rightDepth) + 1
	}

	var result int
	dfs(root, &result)

	return result
}
