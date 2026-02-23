typealias CharFrequencies = Map<Char, Int>

class Solution {
    fun groupAnagrams(strs: Array<String>): List<List<String>> {
        val groups: MutableMap<CharFrequencies, MutableList<String>> = HashMap()
        for (str in strs) {
            val freqs = charFrequencies(str)

            val group = groups.getOrDefault(freqs, mutableListOf())
            group.add(str)

            groups[freqs] = group
        }

        return groups.values.toList()
    }

    fun charFrequencies(str: String): CharFrequencies {
        val chars: MutableMap<Char, Int> = HashMap(str.length)
        for (char in str) {
            chars[char] = chars.getOrDefault(char, 0) + 1
        }

        return chars
    }
}

val res = Solution().groupAnagrams(arrayOf("eat", "tea", "tan", "ate", "nat", "bat"))
println(res)
