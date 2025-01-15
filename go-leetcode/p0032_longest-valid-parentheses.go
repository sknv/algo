package leetcode

// https://leetcode.com/problems/longest-valid-parentheses/description/

func longestValidParentheses(s string) int {
	stack := NewDeque[int](len(s))
	stack.PushBack(-1)

	var maxLen int

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack.PushBack(i)
		} else {
			_, _ = stack.PopBack()

			if stack.IsEmpty() {
				stack.PushBack(i)
			} else {
				lastIdx, _ := stack.PeekBack()
				curLen := i - lastIdx

				maxLen = max(maxLen, curLen)
			}
		}
	}

	return maxLen
}
