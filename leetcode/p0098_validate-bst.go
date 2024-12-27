package leetcode

import "math"

// https://leetcode.com/problems/validate-binary-search-tree/description/

func isValidBST(root *IntTreeNode) bool {
	var dfs func(node *IntTreeNode, min int, max int) bool
	dfs = func(node *IntTreeNode, min, max int) bool {
		if node == nil {
			return true
		}

		if node.Val <= min || node.Val >= max {
			return false
		}

		leftSubtree := dfs(node.Left, min, node.Val)
		rightSubtree := dfs(node.Right, node.Val, max)

		return leftSubtree && rightSubtree
	}

	return dfs(root, math.MinInt64, math.MaxInt64)
}
