package queue

import (
	"container/list"
)

type Queue struct {
	list *list.List
}

func New() *Queue {
	return &Queue{
		list: list.New(),
	}
}

func (q *Queue) IsEmpty() bool {
	return q.list.Front() == nil
}

func (q *Queue) Enqueue(value interface{}) {
	q.list.PushBack(value)
}

func (q *Queue) Dequeue() interface{} {
	if el := q.list.Front(); el != nil {
		q.list.Remove(el)
		return el.Value
	}
	return nil
}

func (q *Queue) Peek() interface{} {
	if el := q.list.Front(); el != nil {
		return el.Value
	}
	return nil
}
