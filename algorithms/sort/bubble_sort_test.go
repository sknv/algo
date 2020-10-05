package sort

import (
	"reflect"
	"testing"
)

func TestBubbleSortInt(t *testing.T) {
	expected, nums := []int{-2, 0, 5, 9, 12}, []int{12, 0, 9, -2, 5}
	BubbleSortInt(nums, func(prev, next int) bool {
		return prev < next
	})
	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("expect BubbleSort(...) to be [%v], got instead [%v]", expected, nums)
	}

	expected, nums = []int{10, 1, -2, -6}, []int{1, -2, 10, -6}
	BubbleSortInt(nums, func(prev, next int) bool {
		return prev > next
	})
	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("expect BubbleSort(...) to be [%v], got instead [%v]", expected, nums)
	}
}
