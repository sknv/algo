package leetcode

// https://leetcode.com/problems/add-two-numbers/description/

func addTwoNumbers(l1 *IntListNode, l2 *IntListNode) *IntListNode {
	preHead := &IntListNode{}
	curNode := preHead

	var carry int

	for l1 != nil || l2 != nil || carry > 0 {
		var val1 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		var val2 int
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		sum := val1 + val2 + carry
		carry = 0
		if sum > 9 {
			sum -= 10
			carry = 1
		}

		curNode.Next = &IntListNode{
			Val:  sum,
			Next: nil,
		}
		curNode = curNode.Next
	}

	return preHead.Next
}
