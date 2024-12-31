package leetcode

import "strconv"

// https://leetcode.com/problems/evaluate-reverse-polish-notation/description/

func evalRPN(tokens []string) int {
	var (
		result int
		stack  = NewDeque[int](len(tokens))
	)

	for _, token := range tokens {
		switch token {
		case "+":
			right, _ := stack.PopBack()
			left, _ := stack.PopBack()

			result = left + right
			stack.PushBack(result)
		case "-":
			right, _ := stack.PopBack()
			left, _ := stack.PopBack()

			result = left - right
			stack.PushBack(result)
		case "*":
			right, _ := stack.PopBack()
			left, _ := stack.PopBack()

			result = left * right
			stack.PushBack(result)
		case "/":
			right, _ := stack.PopBack()
			left, _ := stack.PopBack()

			result = left / right
			stack.PushBack(result)
		default: // number
			result, _ = strconv.Atoi(token)
			stack.PushBack(result)
		}
	}

	return result
}
