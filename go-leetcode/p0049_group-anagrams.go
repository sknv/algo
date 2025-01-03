package leetcode

// https://leetcode.com/problems/group-anagrams/description/

func groupAnagrams(strs []string) [][]string {
	// Группируем по кол-во вхождений каждого символа в текущую строку.
	groupsMap := make(map[[26]byte][]string, len(strs))
	for _, str := range strs {
		freqs := letterFrequencesForAnagrams(str)
		groupsMap[freqs] = append(groupsMap[freqs], str)
	}

	groups := make([][]string, 0, len(groupsMap))
	for _, group := range groupsMap {
		groups = append(groups, group)
	}

	return groups
}

func letterFrequencesForAnagrams(str string) [26]byte {
	var freqs [26]byte
	for i := 0; i < len(str); i++ {
		letter := str[i] - 'a'
		freqs[letter] += 1
	}

	return freqs
}
