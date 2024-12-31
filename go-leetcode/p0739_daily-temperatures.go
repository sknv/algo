package leetcode

// https://leetcode.com/problems/daily-temperatures/description/

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures)) // stack of indexes

	for i, temp := range temperatures {
		for len(stack) > 0 && temp > temperatures[stack[len(stack)-1]] {
			lastIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			result[lastIndex] = i - lastIndex
		}

		stack = append(stack, i)
	}

	return result
}
