package queue

import (
	"testing"
)

func TestQueue_IsEmpty(t *testing.T) {
	queue := New()
	if !queue.IsEmpty() {
		t.Error("expect Queue to be empty")
	}
}

func TestQueue_Enqueue(t *testing.T) {
	queue := New()

	queue.Enqueue(1)
	if queue.IsEmpty() {
		t.Error("expect Queue to have some values")
	}
}

func TestQueue_Dequeue(t *testing.T) {
	queue := New()

	queue.Enqueue(2)
	queue.Enqueue(3)

	if expected, val := 2, queue.Dequeue(); expected != val {
		t.Errorf("expect enqueued value to be [%d], got instead [%d]", expected, val)
	}

	if expected, val := 3, queue.Dequeue(); expected != val {
		t.Errorf("expect enqueued value to be [%d], got instead [%d]", expected, val)
	}

	if !queue.IsEmpty() {
		t.Error("expect Queue to be empty")
	}
}

func TestQueue_Peek(t *testing.T) {
	queue := New()

	queue.Enqueue(2)
	queue.Enqueue(3)

	if expected, val := 2, queue.Peek(); expected != val {
		t.Errorf("expect enqueued value to be [%d], got instead [%d]", expected, val)
	}

	if expected, val := 2, queue.Peek(); expected != val {
		t.Errorf("expect enqueued value to be [%d], got instead [%d]", expected, val)
	}
}
