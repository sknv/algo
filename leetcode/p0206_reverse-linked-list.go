package leetcode

// https://leetcode.com/problems/reverse-linked-list/description/

func reverseList(head *IntListNode) *IntListNode {
	var prev *IntListNode

	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}

	return prev
}
