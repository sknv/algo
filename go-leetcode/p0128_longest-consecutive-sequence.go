package leetcode

// https://leetcode.com/problems/longest-consecutive-sequence/description/

func longestConsecutive(nums []int) int {
	seen := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		seen[num] = struct{}{}
	}

	// Для каждого элемента ищем элемент, меньший на 1.
	// Если не нашли, то для текущего элемента последовательно считаем кол-во больших на 1 элементов.
	var longest int

	for num := range seen {
		if _, found := seen[num-1]; found {
			continue
		}

		curLongest := 1
		for {
			if _, found := seen[num+curLongest]; found {
				curLongest++
			} else {
				break
			}
		}

		longest = max(longest, curLongest)
	}

	return longest
}
