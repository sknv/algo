class Solution {
    fun threeSum(nums: IntArray): List<List<Int>> {
        nums.sort()

        val triplets = hashSetOf<Triple<Int, Int, Int>>()

        for (i in nums.indices) {
            // Если справа остались только положительные элементы, то их сумма не может быть равна 0
            if (nums[i] > 0) {
                break
            }

            val rest = nums.sliceArray(i + 1 until nums.size)
            val target = -nums[i] // сумма двух остальных элементов должна быть равна обратному значению текущего

            val pairs = allTwoSumForSorted(rest, target)
            for (pair in pairs) {
                val triplet = Triple(nums[i], pair.first, pair.second)
                triplets.add(triplet)
            }
        }

        return triplets.map { it.toList() }
    }

    fun allTwoSumForSorted(nums: IntArray, target: Int): List<Pair<Int, Int>> {
        val pairs = mutableListOf<Pair<Int, Int>>()

        var left = 0
        var right = nums.size - 1

        while (left < right) {
            val sum = nums[left] + nums[right]
            when {
                sum < target -> left++
                sum > target -> right--
                else -> {
                    pairs.add(Pair(nums[left], nums[right]))

                    // Сдвигаем левый указатель вправо до тех пор, пока не встретим отличающийся элемент
                    // Это позволяет нам также сдвинуть правый указатель влево, так как для текущего мы уже нашли пару
                    left++
                    right--
                    while (left < right && nums[left] == nums[left - 1]) {
                        left++
                    }
                }
            }
        }

        return pairs
    }
}

val res = Solution().threeSum(intArrayOf(-1, 0, 1, 2, -1, -4))
println(res)
