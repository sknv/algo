package leetcode

// https://leetcode.com/problems/word-subsets/description/

func wordSubsets(words1 []string, words2 []string) []string {
	// Группируем по кол-ву вхождений каждого символа в текущую строку.
	freqs2 := make([][26]byte, 0, len(words2))
	for _, word2 := range words2 {
		freqs2 = append(freqs2, letterFrequencesForWordSubsets(word2))
	}

	// Один раз считаем общее кол-во вхождений символов в words2.
	freqs2Total := reduceFrequencesForWordSubsets(freqs2)

	var result []string

	// Проверяем, входит ли общая частота символов в строки первого слайса.
	for _, word1 := range words1 {
		freq1 := letterFrequencesForWordSubsets(word1)

		if containsWordSubset(freq1, freqs2Total) {
			result = append(result, word1)
		}
	}

	return result
}

func letterFrequencesForWordSubsets(str string) [26]byte {
	var freqs [26]byte
	for i := 0; i < len(str); i++ {
		letter := str[i] - 'a'
		freqs[letter] += 1
	}

	return freqs
}

func reduceFrequencesForWordSubsets(freqs [][26]byte) [26]byte {
	var result [26]byte
	for i := 0; i < len(freqs); i++ {
		freq := freqs[i]

		for j := 0; j < len(freq); j++ {
			result[j] = max(result[j], freq[j])
		}
	}

	return result
}

func containsWordSubset(freq1 [26]byte, freq2 [26]byte) bool {
	for i := 0; i < len(freq1); i++ {
		if freq2[i] > freq1[i] {
			return false
		}
	}

	return true
}
