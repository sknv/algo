class Solution {
    fun twoSum(numbers: IntArray, target: Int): IntArray {
        var left = 0
        var right = numbers.size - 1

        while (left < right) {
            val sum = numbers[left] + numbers[right]
            when {
                sum < target -> left++
                sum > target -> right--
                else -> return intArrayOf(left + 1, right + 1) // ожидается результат для массива с индексом от 1
            }
        }

        throw IllegalArgumentException("No solution")
    }
}

val res = Solution().twoSum(intArrayOf(2, 7, 11, 15), 9)
println(res.contentToString())
