package math

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	if expected, fib := 1, Fibonacci(1); fib != expected {
		t.Errorf("expect Fibonacci(...) to be [%d], got instead [%d]", expected, fib)
	}

	if expected, fib := 1, Fibonacci(2); fib != expected {
		t.Errorf("expect Fibonacci(...) to be [%d], got instead [%d]", expected, fib)
	}

	if expected, fib := 8, Fibonacci(6); fib != expected {
		t.Errorf("expect Fibonacci(...) to be [%d], got instead [%d]", expected, fib)
	}
}
