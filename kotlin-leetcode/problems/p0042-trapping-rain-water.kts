import kotlin.math.max

class Solution {
    fun trap(height: IntArray): Int {
        if (height.isEmpty()) {
            return 0
        }

        var left = 0
        var right = height.size - 1

        var maxLeft = height[left]
        var maxRight = height[right]

        var result = 0

        while (left < right) {
            // Идея в том, чтобы с помощью двух указателей для каждой клетки собрать объем воды либо слева, либо справа
            // в зависимости от того, какая из высот в данный момент меньше
            if (maxLeft < maxRight) {
                left++

                // Собираем объем воды слева от текущей позиции
                val curHeight = height[left]
                maxLeft = max(maxLeft, curHeight)
                val curVol = maxLeft - curHeight
                result += curVol
            } else {
                right--

                // Собираем объем воды справа от текущей позиции
                val curHeight = height[right]
                maxRight = max(maxRight, curHeight)
                val curVol = maxRight - curHeight
                result += curVol
            }
        }

        return result
    }
}

val res = Solution().trap(intArrayOf(0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1))
println(res)
