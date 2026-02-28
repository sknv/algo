class Solution {
    fun productExceptSelf(nums: IntArray): IntArray {
        val result = IntArray(nums.size)
        result[0] = 1

        // Собираем произведения с левой стороны
        for (i in 1 until nums.size) {
            result[i] = result[i-1] * nums[i-1]
        }

        // Собираем произведения с правой стороны
        var acc = 1
        for (i in nums.size - 1 downTo 0) {
            result[i] *= acc
            acc *= nums[i]
        }

        return result
    }
}

val res = Solution().productExceptSelf(intArrayOf(1, 2, 3, 4))
println(res.contentToString())
