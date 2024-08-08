package leetcode

//
// Linked lists
//

type IntListNode struct {
	Val  int
	Next *IntListNode
}

//
// Deque
//

type Deque[T any] struct {
	data []T
}

func NewDeque[T any](capacity int) Deque[T] {
	return Deque[T]{
		data: make([]T, 0, capacity),
	}
}

// PushFront adds an element to the front of the deque.
func (d *Deque[T]) PushFront(value T) {
	d.data = append([]T{value}, d.data...)
}

// PushBack adds an element to the back of the deque.
func (d *Deque[T]) PushBack(value T) {
	d.data = append(d.data, value)
}

// PopFront removes and returns an element from the front of the deque.
func (d *Deque[T]) PopFront() (T, bool) {
	if len(d.data) == 0 {
		var zero T

		return zero, false
	}

	value := d.data[0]
	d.data = d.data[1:]

	return value, true
}

// PopBack removes and returns an element from the back of the deque.
func (d *Deque[T]) PopBack() (T, bool) {
	if len(d.data) == 0 {
		var zero T

		return zero, false
	}

	value := d.data[len(d.data)-1]
	d.data = d.data[:len(d.data)-1]

	return value, true
}
