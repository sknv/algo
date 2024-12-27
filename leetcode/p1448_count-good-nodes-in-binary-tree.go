package leetcode

// https://leetcode.com/problems/count-good-nodes-in-binary-tree/description/

func goodNodes(root *IntTreeNode) int {
	var dfs func(node *IntTreeNode, maxValue int) int
	dfs = func(node *IntTreeNode, maxValue int) int {
		if node == nil {
			return 0
		}

		var result int
		if node.Val >= maxValue {
			result = 1
		}

		maxValue = max(maxValue, node.Val)

		result += dfs(node.Left, maxValue)
		result += dfs(node.Right, maxValue)

		return result
	}

	return dfs(root, root.Val)
}
