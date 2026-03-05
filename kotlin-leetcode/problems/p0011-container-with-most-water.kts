import kotlin.math.max
import kotlin.math.min

class Solution {
    fun maxArea(height: IntArray): Int {
        var left = 0
        var right = height.size - 1
        var maxWater = 0

        while (left < right) {
            val leftHeight = height[left]
            val rightHeight = height[right]
            val curWater = (right - left) * min(leftHeight, rightHeight)
            maxWater = max(maxWater, curWater)

            if (leftHeight < rightHeight) {
                left++
            } else {
                right--
            }
        }

        return maxWater
    }
}

val res = Solution().maxArea(intArrayOf(1, 8, 6, 2, 5, 4, 8, 3, 7))
println(res)
