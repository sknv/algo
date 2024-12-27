package leetcode

// https://leetcode.com/problems/binary-tree-right-side-view/description/

func rightSideView(root *IntTreeNode) []int {
	var result []int

	queue := NewDeque[*IntTreeNode](0)
	if root != nil {
		queue.PushBack(root)
	}

	for !queue.IsEmpty() {
		var rightValue int
		curLen := queue.Len()

		for i := 0; i < curLen; i++ {
			node, _ := queue.PopFront()
			rightValue = node.Val

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		result = append(result, rightValue)
	}

	return result
}
