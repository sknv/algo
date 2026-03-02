class Solution {
    fun isPalindrome(s: String): Boolean {
        var left = 0
        var right = s.length - 1

        while (left < right) {
            // Отсекаем все лишние символы
            // Не забываем проверки на выход за границы строки, но чуть более оптимизированные
            while (left < right && !s[left].isLetterOrDigit()) {
                left++
            }

            while (right > left && !s[right].isLetterOrDigit()) {
                right--
            }

            if (s[left].lowercaseChar() != s[right].lowercaseChar()) {
                return false
            }

            left++
            right--
        }

        return true
    }
}

val res = Solution().isPalindrome("A man, a plan, a canal: Panama ")
println(res)
