package leetcode

// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/

func maxDepth(root *IntTreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	return max(leftDepth, rightDepth) + 1
}
