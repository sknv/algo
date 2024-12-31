package leetcode

// https://leetcode.com/problems/combination-sum/description/

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	backtrackCombinationSum(candidates, target, 0, []int{}, &result)

	return result
}

func backtrackCombinationSum(candidates []int, target int, start int, current []int, result *[][]int) {
	// If the target is reached, add the current combination to the result
	if target == 0 {
		newCombination := make([]int, len(current))
		copy(newCombination, current)
		*result = append(*result, newCombination)

		return
	}

	// Iterate through the candidates starting from the given index
	for i := start; i < len(candidates); i++ {
		// Skip candidates that are larger than the remaining target
		if candidates[i] > target {
			continue
		}

		// Include the current candidate in the combination
		current = append(current, candidates[i])
		// Recurse with the updated target and the same start index (to allow reuse)
		backtrackCombinationSum(candidates, target-candidates[i], i, current, result)
		// Backtrack by removing the last candidate
		current = current[:len(current)-1]
	}
}
