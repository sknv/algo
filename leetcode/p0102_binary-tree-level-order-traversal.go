package leetcode

// https://leetcode.com/problems/binary-tree-level-order-traversal/description/

func levelOrder(root *IntTreeNode) [][]int {
	var result [][]int

	queue := NewDeque[*IntTreeNode](0)
	if root != nil {
		queue.PushBack(root)
	}

	for !queue.IsEmpty() {
		var curLevel []int
		curLen := queue.Len()

		for i := 0; i < curLen; i++ {
			node, _ := queue.PopFront()
			curLevel = append(curLevel, node.Val)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}

			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}

		result = append(result, curLevel)
	}

	return result
}
