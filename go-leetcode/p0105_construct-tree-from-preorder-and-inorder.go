package leetcode

// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/

func buildTreeFromPreorderAndInorder(preorder []int, inorder []int) *IntTreeNode {
	inorderIndices := make(map[int]int)
	for i, val := range inorder {
		inorderIndices[val] = i
	}

	var curPreorderIndex int

	var dfs func(left int, right int) *IntTreeNode
	dfs = func(left int, right int) *IntTreeNode {
		if left > right {
			return nil
		}

		curValue := preorder[curPreorderIndex]
		curPreorderIndex++

		node := IntTreeNode{
			Val: curValue,
		}

		mid := inorderIndices[curValue]

		node.Left = dfs(left, mid-1)
		node.Right = dfs(mid+1, right)

		return &node
	}

	return dfs(0, len(inorder)-1)
}
