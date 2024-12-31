package leetcode

import (
	"math"
	"slices"
)

// https://leetcode.com/problems/koko-eating-bananas/description/

func minEatingSpeed(piles []int, h int) int {
	// Если предположить, что Коко ест бананы с максимальной скоростью,
	// ей понадобится len(piles) часов. Таким образом, максимальное число в массиве есть наша верхняя граница.
	var (
		left     = 1 // Начинаем с 1 банана в час
		right    = slices.Max(piles)
		minSpeed int
	)

	for left <= right {
		speed := left + (right-left)/2

		// Считаем кол-во времени, которое понадобится для того, чтобы съесть все бананы
		if totalTime := totalEatingTime(piles, speed); totalTime <= h {
			minSpeed = speed
			right = speed - 1
		} else {
			left = speed + 1
		}
	}

	return minSpeed
}

func totalEatingTime(piles []int, speed int) int {
	var totalTime int
	for _, bananas := range piles {
		totalTime += int(math.Ceil(float64(bananas) / float64(speed)))
	}

	return totalTime
}
