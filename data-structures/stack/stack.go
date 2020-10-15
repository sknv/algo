package stack

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

func New() *Stack {
	return &Stack{
		list: list.New(),
	}
}

func (s *Stack) IsEmpty() bool {
	return s.list.Front() == nil
}

func (s *Stack) Push(value interface{}) {
	s.list.PushFront(value)
}

func (s *Stack) Pop() interface{} {
	if el := s.list.Front(); el != nil {
		s.list.Remove(el)
		return el.Value
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	if el := s.list.Front(); el != nil {
		return el.Value
	}
	return nil
}
