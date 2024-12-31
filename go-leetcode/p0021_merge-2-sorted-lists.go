package leetcode

// https://leetcode.com/problems/reverse-linked-list/description/

func mergeTwoLists(list1 *IntListNode, list2 *IntListNode) *IntListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)

		return list1
	}

	list2.Next = mergeTwoLists(list1, list2.Next)

	return list2
}
