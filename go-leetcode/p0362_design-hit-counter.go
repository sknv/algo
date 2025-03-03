package leetcode

import (
	"slices"
)

// https://leetcode.com/problems/design-hit-counter/description/

type HitCounter struct {
	timestamps []int
}

func Constructor() HitCounter {
	return HitCounter{}
}

func (this *HitCounter) Hit(timestamp int) {
	this.timestamps = append(this.timestamps, timestamp)
}

func (this *HitCounter) GetHits(timestamp int) int {
	last300secIdx, _ := slices.BinarySearch(this.timestamps, timestamp-300+1)

	return len(this.timestamps) - last300secIdx
}
