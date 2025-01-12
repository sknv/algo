package leetcode

// https://leetcode.com/problems/check-if-a-parentheses-string-can-be-valid/description/

func canBeValid(s string, locked string) bool {
	length := len(s)
	if length%2 == 1 {
		return false
	}

	openBrackets := NewDeque[int](length)
	unlocked := NewDeque[int](length)

	// Проходимся по строке и пытаемся сверить кол-во закрывающихся скобок
	// с кол-вом открывающихся (в приоритете) или же с теми, которые можо изменить.
	for i := 0; i < length; i++ {
		if locked[i] == '0' {
			unlocked.PushBack(i)
		} else if s[i] == '(' {
			openBrackets.PushBack(i)
		} else if s[i] == ')' {
			if !openBrackets.IsEmpty() {
				_, _ = openBrackets.PopBack()
			} else if !unlocked.IsEmpty() {
				_, _ = unlocked.PopBack()
			} else {
				return false
			}
		}
	}

	// Если у нас остались октрывающиеся скобки и те позиции, которые можно изменить, то проверяем, что,
	// во-первых, их кол-во открывающихся меньше либо равно кол-во потенциально измененных, а,
	// во-вторых, позиция открыающейся скобки должна быть левее по условию валидности в задаче.
	for !openBrackets.IsEmpty() && !unlocked.IsEmpty() {
		bracketPos, _ := openBrackets.PopBack()
		unlockedPos, _ := unlocked.PopBack()

		if bracketPos > unlockedPos {
			return false
		}
	}

	return openBrackets.IsEmpty()
}
