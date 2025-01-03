package leetcode

// https://leetcode.com/problems/house-robber/description/

func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}

	var dp func(i int) int
	dp = func(i int) int {
		if i >= len(nums) {
			return 0
		}

		if memo[i] != -1 {
			return memo[i]
		}

		curMoney := nums[i] + dp(i+2)
		nextMoney := dp(i + 1)
		memo[i] = max(curMoney, nextMoney)

		return memo[i]
	}

	return dp(0)
}
