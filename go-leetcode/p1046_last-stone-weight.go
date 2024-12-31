package leetcode

import "container/heap"

// https://leetcode.com/problems/last-stone-weight/description/

func lastStoneWeight(stones []int) int {
	pq := make(PriorityQueue[int], 0, len(stones))
	for _, stone := range stones {
		heap.Push(&pq, &PriorityQueueItem[int]{
			Value:    stone,
			Priority: stone,
		})
	}

	for pq.Len() > 1 {
		first := heap.Pop(&pq).(*PriorityQueueItem[int]).Value
		second := heap.Pop(&pq).(*PriorityQueueItem[int]).Value

		if diff := first - second; diff > 0 {
			heap.Push(&pq, &PriorityQueueItem[int]{
				Value:    diff,
				Priority: diff,
			})
		}
	}

	// Если камней не осталось, то добавим нулевое значение.
	heap.Push(&pq, &PriorityQueueItem[int]{
		Value:    0,
		Priority: 0,
	})

	last := heap.Pop(&pq).(*PriorityQueueItem[int]).Value
	return last
}
