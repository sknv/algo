package leetcode

// https://leetcode.com/problems/reverse-nodes-in-k-group/

func reverseKGroup(head *IntListNode, k int) *IntListNode {
	totalElements := totalListElements(head)
	groupCount := totalElements / k

	var (
		curGroup int

		prevGroupHead *IntListNode
		prevGroupNext = head

		newHead *IntListNode
	)

	for curGroup < groupCount {
		originalHead := prevGroupNext
		head, prevGroupNext = reverseKGroupList(originalHead, k)

		if newHead == nil { // Сохраняем начало нового списка.
			newHead = head
		}

		if prevGroupHead != nil { // Соединяем конец предыдушей группы с новой группой.
			prevGroupHead.Next = head
		}
		prevGroupHead = originalHead

		curGroup++
	}

	// Цепляем остаток к последнему элементу.
	prevGroupHead.Next = prevGroupNext

	return newHead
}

func totalListElements(head *IntListNode) int {
	var total int

	for head != nil {
		total++
		head = head.Next
	}

	return total
}

func reverseKGroupList(head *IntListNode, k int) (*IntListNode, *IntListNode) {
	var (
		prev  *IntListNode
		next  *IntListNode
		shift int
	)

	for head != nil && shift < k {
		next = head.Next
		head.Next = prev
		prev, head = head, next

		shift++
	}

	return prev, next
}
