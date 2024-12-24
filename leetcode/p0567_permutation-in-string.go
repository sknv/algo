package leetcode

// https://leetcode.com/problems/permutation-in-string/description/

func checkInclusion(s1 string, s2 string) bool {
	right := len(s1)
	if len(s2) < right {
		return false
	}

	originalFreqs := newInclusionSlidingWindow(s1)
	curFreqs := newInclusionSlidingWindow(s2[:right])
	if originalFreqs.equals(curFreqs) {
		return true
	}

	for right < len(s2) {
		curFreqs.slideRight(s2[right])
		if originalFreqs.equals(curFreqs) {
			return true
		}

		right++
	}

	return false
}

type inclusionSlidingWindow struct {
	value            []byte
	letterFrequences [26]byte
}

func newInclusionSlidingWindow(str string) inclusionSlidingWindow {
	var freqs [26]byte
	for i := 0; i < len(str); i++ {
		letter := str[i] - 'a'
		freqs[letter] += 1
	}

	return inclusionSlidingWindow{
		value:            []byte(str),
		letterFrequences: freqs,
	}
}

func (w *inclusionSlidingWindow) slideRight(rightChar byte) {
	leftChar := w.value[0]
	w.letterFrequences[leftChar-'a']--

	w.value = append(w.value[1:len(w.value)], rightChar)
	w.letterFrequences[rightChar-'a']++
}

func (w inclusionSlidingWindow) equals(other inclusionSlidingWindow) bool {
	return w.letterFrequences == other.letterFrequences
}
