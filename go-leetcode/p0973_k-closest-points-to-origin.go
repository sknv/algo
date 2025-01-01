package leetcode

// https://leetcode.com/problems/k-closest-points-to-origin/description/

func kClosest(points [][]int, k int) [][]int {
	minHeap := NewBinaryHeap(k+1, func(prev, next []int) bool { return prev[0] < next[0] })
	for _, point := range points {
		dist := estimateEuclideanDistanceToZero(point)
		minHeap.Push([]int{dist, point[0], point[1]})
	}

	result := make([][]int, 0, k)
	for i := 0; i < k; i++ {
		nextPoint := minHeap.Pop()
		result = append(result, []int{nextPoint[1], nextPoint[2]})
	}

	return result
}

func estimateEuclideanDistanceToZero(point []int) int {
	x, y := point[0], point[1]

	return x*x + y*y
}
