class Solution {
    fun containsDuplicate(nums: IntArray): Boolean {
        val seen: MutableSet<Int> = HashSet(nums.size)

        for (num in nums) {
            if (seen.contains(num)) {
                return true
            }

            seen.add(num)
        }

        return false
    }
}

val res = Solution().containsDuplicate(intArrayOf(1, 2, 3, 1))
println(res)
