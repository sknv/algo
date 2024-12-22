package leetcode

// https://leetcode.com/problems/group-anagrams/description/

func groupAnagrams(strs []string) [][]string {
	groupsMap := make(map[[26]byte][]string, len(strs))
	for _, str := range strs {
		freqs := letterFrequences(str)
		groupsMap[freqs] = append(groupsMap[freqs], str)
	}

	groups := make([][]string, 0, len(groupsMap))
	for _, group := range groupsMap {
		groups = append(groups, group)
	}

	return groups
}

func letterFrequences(str string) [26]byte {
	var freqs [26]byte
	for i := 0; i < len(str); i++ {
		letter := str[i] - 'a'
		freqs[letter] += 1
	}

	return freqs
}
