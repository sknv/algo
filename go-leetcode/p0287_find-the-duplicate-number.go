package leetcode

// https://leetcode.com/problems/find-the-duplicate-number/description/

func findDuplicate(nums []int) int {
	for _, num := range nums {
		idx := abs(num) - 1
		if element := nums[idx]; element < 0 {
			return abs(num)
		}

		nums[idx] *= -1
	}

	return 0
}
