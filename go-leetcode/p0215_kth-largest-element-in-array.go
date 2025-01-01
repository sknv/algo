package leetcode

// https://leetcode.com/problems/kth-largest-element-in-an-array/description/

func findKthLargest(nums []int, k int) int {
	minHeap := NewMinHeap[int](k + 1)
	for _, num := range nums {
		minHeap.Push(num)

		if minHeap.Len() > k {
			_ = minHeap.Pop()
		}
	}

	return minHeap.Pop()
}
