package leetcode

import "container/heap"

// https://leetcode.com/problems/top-k-frequent-elements/description/

func topKFrequent(nums []int, k int) []int {
	// Считаем как часто встречаются элементы в массиве.
	seen := make(map[int]int, len(nums))
	for _, num := range nums {
		seen[num] += 1
	}

	// Создаем очередь с приоритетом на основе посчитанных вхождений.
	pq := make(PriorityQueue[int], 0, len(seen))
	for value, priority := range seen {
		heap.Push(&pq, &PriorityQueueItem[int]{
			Value:    value,
			Priority: priority,
		})
	}

	// Достаем из очереди k нужных элементов.
	freq := make([]int, 0, k)
	for i := 0; i < k; i++ {
		freq = append(freq, heap.Pop(&pq).(*PriorityQueueItem[int]).Value)
	}

	return freq
}
