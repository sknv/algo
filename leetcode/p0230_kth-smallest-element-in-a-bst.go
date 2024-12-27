package leetcode

// https://leetcode.com/problems/kth-smallest-element-in-a-bst/description/

func kthSmallest(root *IntTreeNode, k int) int {
	var (
		arr        []int
		inOrderDFS func(node *IntTreeNode)
	)

	inOrderDFS = func(node *IntTreeNode) {
		if node == nil {
			return
		}

		inOrderDFS(node.Left)
		arr = append(arr, node.Val)
		inOrderDFS(node.Right)
	}

	inOrderDFS(root)

	return arr[k-1]
}
