package math

func Fibonacci(number int) int {
	cur, next := 0, 1
	for i := 0; i < number; i++ {
		cur, next = next, cur+next
	}
	return cur
}
