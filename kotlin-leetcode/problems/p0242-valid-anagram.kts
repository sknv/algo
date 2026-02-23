class Solution {
    fun isAnagram(s: String, t: String): Boolean {
        if (s.length != t.length) {
            return false
        }

        val seen: MutableMap<Char, Int> = HashMap(s.length)
        for (char in s) {
            seen[char] = seen.getOrDefault(char, 0) + 1
        }

        for (char in t) {
            val charCount = seen.getOrDefault(char, 0) - 1
            if (charCount < 0) {
                return false
            }

            seen[char] = charCount
        }

        return true
    }
}

val res = Solution().isAnagram("anagram", "nagaram")
println(res)
