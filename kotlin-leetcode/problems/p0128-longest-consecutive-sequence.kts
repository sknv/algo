import kotlin.math.max

class Solution {
    fun longestConsecutive(nums: IntArray): Int {
        val numSet = nums.toHashSet()

        var result = 0
        for (num in nums) {
            // Ищем начало последовательности: если есть число, меньшее на 1, то текущий элемент не является началом
            if (numSet.contains(num - 1)) {
                continue
            }

            // Считаем кол-во последовательных элементов и сравниваем с рекордом
            var curCount = 1
            while (numSet.contains(num + curCount)) {
                curCount++
            }

            result = max(result, curCount)
        }

        return result
    }
}

val res = Solution().longestConsecutive(intArrayOf(100, 4, 200, 1, 3, 2))
println(res)
