package search

import (
	"testing"
)

func TestBinarySearchInt(t *testing.T) {
	var nums = []int{-2, 0, 5, 9, 12}

	target := 9
	if idx, expected := BinarySearchInt(nums, target), 3; idx != expected {
		t.Errorf("expect binarySearch(..., %d) to be: %d, got instead: %d", target, expected, idx)
	}

	target = -2
	if idx, expected := BinarySearchInt(nums, target), 0; idx != expected {
		t.Errorf("expect binarySearch(..., %d) to be: %d, got instead: %d", target, expected, idx)
	}

	target = 6
	if idx, expected := BinarySearchInt(nums, target), -1; idx != expected {
		t.Errorf("expect binarySearch(..., %d) to be: %d, got instead: %d", target, expected, idx)
	}
}
