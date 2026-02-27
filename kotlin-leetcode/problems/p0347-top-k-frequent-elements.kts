import java.util.PriorityQueue

class Solution {
    fun topKFrequent(nums: IntArray, k: Int): IntArray {
        // Считаем кол-во вхождений элементов
        val freqs: MutableMap<Int, Int> = HashMap(nums.size)
        for (num in nums) {
            freqs[num] = freqs.getOrDefault(num, 0) + 1
        }

        // Формируем очередь с приоритетом: чем чаще элемент встречается в массиве, тем больше у него приоритет
        val compareByValue: Comparator<Map.Entry<Int, Int>> = compareByDescending { it.value }
        val pq = PriorityQueue(compareByValue)
        for (pair in freqs.entries) {
            pq.add(pair)
        }

        // Забираем k элементов из очереди с приоритетом
        val result: MutableList<Int> = ArrayList(k)
        repeat(k) {
            val top = pq.remove()
            result.add(top.key)
        }

        return result.toIntArray()
    }
}

val res = Solution().topKFrequent(intArrayOf(1, 1, 1, 2, 2, 3, 3, 3, 3), 2)
println(res.contentToString())
