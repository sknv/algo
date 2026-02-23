class Solution {
    fun twoSum(nums: IntArray, target: Int): IntArray {
        val seen: MutableMap<Int, Int> = HashMap(nums.size)

        for ((idx, num) in nums.withIndex()) {
            val search = target - num

            val searchIdx = seen[search]
            if (searchIdx != null) {
                return intArrayOf(searchIdx, idx)
            }

            seen[num] = idx
        }

        throw IllegalArgumentException("No solution")
    }
}

val res = Solution().twoSum(intArrayOf(2, 7, 11, 15), 9)
println(res.contentToString())
