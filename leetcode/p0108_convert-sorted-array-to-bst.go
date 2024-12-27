package leetcode

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/description/

func sortedArrayToBST(nums []int) *IntTreeNode {
	return buildTree(nums, 0, len(nums)-1)
}

func buildTree(nums []int, left, right int) *IntTreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	node := &IntTreeNode{Val: nums[mid]}
	node.Left = buildTree(nums, left, mid-1)
	node.Right = buildTree(nums, mid+1, right)

	return node
}
