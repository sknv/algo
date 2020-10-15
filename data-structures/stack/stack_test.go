package stack

import (
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	stack := New()
	if !stack.IsEmpty() {
		t.Error("expect Stack to be empty")
	}
}

func TestStack_Enqueue(t *testing.T) {
	stack := New()

	stack.Push(1)
	if stack.IsEmpty() {
		t.Error("expect Stack to have some values")
	}
}

func TestStack_Dequeue(t *testing.T) {
	stack := New()

	stack.Push(2)
	stack.Push(3)

	if expected, val := 3, stack.Pop(); expected != val {
		t.Errorf("expect pushed value to be [%d], got instead [%d]", expected, val)
	}

	if expected, val := 2, stack.Pop(); expected != val {
		t.Errorf("expect pushed value to be [%d], got instead [%d]", expected, val)
	}

	if !stack.IsEmpty() {
		t.Error("expect Stack to be empty")
	}
}

func TestStack_Peek(t *testing.T) {
	stack := New()

	stack.Push(2)
	stack.Push(3)

	if expected, val := 3, stack.Peek(); expected != val {
		t.Errorf("expect pushed value to be [%d], got instead [%d]", expected, val)
	}

	if expected, val := 3, stack.Peek(); expected != val {
		t.Errorf("expect pushed value to be [%d], got instead [%d]", expected, val)
	}
}
