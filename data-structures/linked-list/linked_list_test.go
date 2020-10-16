package linked_list

import (
	"reflect"
	"testing"
)

func TestLinkedList_Reverse(t *testing.T) {
	list := makeLinkedList()
	list.Reverse()

	var values []int
	list.Each(func(value interface{}) {
		values = append(values, value.(int))
	})

	expected := []int{4, 3, 2, 1, 0}
	if !reflect.DeepEqual(expected, values) {
		t.Errorf("expect reversed values to be [%v], got instead [%v]", expected, values)
	}
}

func makeLinkedList() *LinkedList {
	zero := LinkedListNode{Value: 0}
	one := LinkedListNode{Value: 1}
	two := LinkedListNode{Value: 2}
	three := LinkedListNode{Value: 3}
	four := LinkedListNode{Value: 4}

	zero.Next = &one
	one.Next = &two
	two.Next = &three
	three.Next = &four
	return &LinkedList{
		Head: &zero,
	}
}
