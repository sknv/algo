class Solution {
    fun isAnagram(s: String, t: String): Boolean {
        if (s.length != t.length) {
            return false
        }

        // Считаем кол-во вхождений символов в первую строку
        val seen: MutableMap<Char, Int> = HashMap(s.length)
        for (char in s) {
            seen[char] = seen.getOrDefault(char, 0) + 1
        }

        // Проверяем, что все символы второй строки встречаются в ней ровно столько же раз, сколько в первой
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
