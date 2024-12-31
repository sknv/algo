package leetcode

// https://leetcode.com/problems/binary-tree-maximum-path-sum/description

func maxPathSum(root *IntTreeNode) int {
	// Initialize a variable to keep track of the maximum path sum globally.
	maxSum := root.Val

	// Helper function to calculate the maximum gain from each node.
	var dfs func(node *IntTreeNode) int
	dfs = func(node *IntTreeNode) int {
		if node == nil {
			return 0
		}

		// Recursively calculate the maximum gain from left and right subtrees.
		leftGain := max(0, dfs(node.Left))
		rightGain := max(0, dfs(node.Right))

		// Calculate the price of a new path that passes through the current node.
		newPathSum := node.Val + leftGain + rightGain

		// Update the global maximum path sum if the new path is better.
		if newPathSum > maxSum {
			maxSum = newPathSum
		}

		// Return the maximum gain if we continue the same path including the current node.
		return node.Val + max(leftGain, rightGain)
	}

	// Start the recursion from the root node.
	dfs(root)

	return maxSum
}
