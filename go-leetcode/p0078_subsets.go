package leetcode

// https://leetcode.com/problems/subsets/description/

func subsets(nums []int) [][]int {
	result := [][]int{}
	backtrackSubsets(nums, 0, []int{}, &result)

	return result
}

func backtrackSubsets(nums []int, start int, current []int, result *[][]int) {
	// Add the current subset to the result
	newSubset := make([]int, len(current))
	copy(newSubset, current)
	*result = append(*result, newSubset)

	// Iterate through the remaining elements
	for i := start; i < len(nums); i++ {
		// Include nums[i] in the current subset
		current = append(current, nums[i])
		// Recurse with the next index
		backtrackSubsets(nums, i+1, current, result)
		// Exclude nums[i] from the current subset (backtrack)
		current = current[:len(current)-1]
	}
}
