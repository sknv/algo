package leetcode

// https://leetcode.com/problems/reverse-linked-list/description/

func reorderList(head *IntListNode) {
	if head.Next == nil {
		return
	}

	middle := middleOfReorderedList(head)        // Находим середину списка.
	reversedHalf := reverseReorderedList(middle) // Разворачиваем.

	mergeReorderedLists(head, reversedHalf) // Соединяем поочередно.
}

func middleOfReorderedList(head *IntListNode) *IntListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverseReorderedList(head *IntListNode) *IntListNode {
	var prev *IntListNode

	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}

	return prev
}

func mergeReorderedLists(head1 *IntListNode, head2 *IntListNode) {
	for head1.Next != nil { // head2 всегда либо такой же либо длиннее, чем head1.
		next1 := head1.Next
		head1.Next = head2
		head1 = next1

		next2 := head2.Next
		head2.Next = head1
		head2 = next2
	}
}
