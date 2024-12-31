package leetcode

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

func lengthOfLongestSubstring(s string) int {
	var left, right int
	var longestLen int

	window := newLongestSubstringWindow()

	for right < len(s) {
		curChar := s[right]
		if window.contains(curChar) {
			longestLen = max(longestLen, right-left)
			window.popFront()
			left++

			continue
		}

		window.pushBack(curChar)
		right++
	}

	return max(longestLen, right-left)
}

type longestSubstringWindow struct {
	chars map[byte]struct{}
	deque Deque[byte]
}

func newLongestSubstringWindow() longestSubstringWindow {
	return longestSubstringWindow{
		chars: make(map[byte]struct{}),
		deque: NewDeque[byte](0),
	}
}

func (w *longestSubstringWindow) contains(char byte) bool {
	_, ok := w.chars[char]

	return ok
}

func (w *longestSubstringWindow) pushBack(char byte) {
	w.deque.PushBack(char)
	w.chars[char] = struct{}{}
}

func (w *longestSubstringWindow) popFront() {
	char, _ := w.deque.PopFront()
	delete(w.chars, char)
}
