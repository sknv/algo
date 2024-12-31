package leetcode

// https://leetcode.com/problems/permutations/description/

func permute(nums []int) [][]int {
	permutations := [][]int{}
	seen := make([]bool, len(nums))
	backtrackPermutations(nums, []int{}, seen, &permutations)

	return permutations
}

func backtrackPermutations(nums []int, current []int, seen []bool, result *[][]int) {
	if len(current) == len(nums) {
		newPermutation := make([]int, len(current))
		copy(newPermutation, current)
		*result = append(*result, newPermutation)

		return
	}

	for i := 0; i < len(nums); i++ {
		if seen[i] {
			continue
		}

		current = append(current, nums[i])
		seen[i] = true

		backtrackPermutations(nums, current, seen, result)

		current = current[:len(current)-1]
		seen[i] = false
	}
}
