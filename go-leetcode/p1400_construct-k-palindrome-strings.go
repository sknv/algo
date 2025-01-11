package leetcode

// https://leetcode.com/problems/construct-k-palindrome-strings/description

func canConstructPalindrome(s string, k int) bool {
	if len(s) < k {
		return false
	}

	// Т.к. нам надо поместить буквы с нечетным кол-вом в k палиндромов,
	// то если таковых больше, чем k, то мы не сможем это сделать.
	freqs := letterFrequencesForConstructingPalindrome(s)

	var oddFreqCount int
	for _, freq := range freqs {
		if freq%2 == 1 {
			oddFreqCount++
		}
	}

	return oddFreqCount <= k
}

func letterFrequencesForConstructingPalindrome(str string) [26]byte {
	var freqs [26]byte
	for i := 0; i < len(str); i++ {
		letter := str[i] - 'a'
		freqs[letter] += 1
	}

	return freqs
}
