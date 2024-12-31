package leetcode

// https://leetcode.com/problems/two-sum/description/

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))

	for curIdx, num := range nums {
		rem := target - num
		if targetIdx, ok := seen[rem]; ok {
			return []int{targetIdx, curIdx}
		}

		seen[num] = curIdx
	}

	return nil
}
