package linked_list

type LinkedListNode struct {
	Value interface{}
	Next  *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
}

func (l *LinkedList) Each(visit func(value interface{})) {
	cur := l.Head
	for cur != nil {
		visit(cur.Value)
		cur = cur.Next
	}
}

func (l *LinkedList) Reverse() {
	var prev *LinkedListNode
	cur := l.Head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	l.Head = prev
}
