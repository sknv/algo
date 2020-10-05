package sort

import (
	"reflect"
	"testing"
)

func TestInsertionSortInt(t *testing.T) {
	expected, nums := []int{-2, 0, 5, 9, 12}, []int{12, 0, 9, -2, 5}
	InsertionSortInt(nums, func(prev, next int) bool {
		return prev < next
	})
	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("expect InsertionSort(...) to be [%v], got instead [%v]", expected, nums)
	}

	expected, nums = []int{10, 1, -2, -6}, []int{1, -2, 10, -6}
	InsertionSortInt(nums, func(prev, next int) bool {
		return prev > next
	})
	if !reflect.DeepEqual(expected, nums) {
		t.Errorf("expect InsertionSort(...) to be [%v], got instead [%v]", expected, nums)
	}
}
