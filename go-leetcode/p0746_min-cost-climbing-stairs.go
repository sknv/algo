package leetcode

// https://leetcode.com/problems/min-cost-climbing-stairs/description/

func minCostClimbingStairs(cost []int) int {
	memo := make([]int, len(cost))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}

	var dp func(i int) int
	dp = func(i int) int {
		if i >= len(cost) {
			return 0
		}

		if memo[i] != -1 {
			return memo[i]
		}

		climb1, climb2 := dp(i+1), dp(i+2)
		memo[i] = cost[i] + min(climb1, climb2)

		return memo[i]
	}

	step0, step1 := dp(0), dp(1)

	return min(step0, step1)
}
