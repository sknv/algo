package leetcode

import "container/heap"

//
// Linked lists
//

type IntListNode struct {
	Val  int
	Next *IntListNode
}

//
// Tree
//

type IntTreeNode struct {
	Val   int
	Left  *IntTreeNode
	Right *IntTreeNode
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
//
// No need to use the method to implement a stack or a queue.
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

//
// Priority queue
//

// An Item is something we manage in a priority queue.
type PriorityQueueItem[T any] struct {
	Value    T   // The value of the item; arbitrary.
	Priority int // The priority of the item in the queue.

	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue[T any] []*PriorityQueueItem[T]

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*PriorityQueueItem[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]

	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue[T]) Update(item *PriorityQueueItem[T], value T, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.index)
}
