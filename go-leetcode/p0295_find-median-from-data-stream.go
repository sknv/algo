package leetcode

import "slices"

// https://leetcode.com/problems/find-median-from-data-stream/description/

// TODO: this is a hard problem and shoud be implemented using max and min heaps for optimal solution.

type MedianFinder struct {
	arr []int
}

func MedianFinderConstructor() MedianFinder {
	return MedianFinder{
		arr: make([]int, 0, 1),
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.arr = append(this.arr, num)
}

func (this *MedianFinder) FindMedian() float64 {
	slices.Sort(this.arr)

	size := len(this.arr)
	if size%2 == 0 {
		return float64(this.arr[size/2-1]+this.arr[size/2]) / 2.0
	}

	return float64(this.arr[size/2])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
