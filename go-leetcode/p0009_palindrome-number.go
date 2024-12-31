package leetcode

// https://leetcode.com/problems/palindrome-number/description/

func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}

	reversed := reverseInt(x)

	return x == reversed
}

func reverseInt(x int) int {
	var reversed int

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return reversed
}
