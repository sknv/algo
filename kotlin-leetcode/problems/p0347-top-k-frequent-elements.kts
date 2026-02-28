import java.util.PriorityQueue

class Solution {
    fun topKFrequent(nums: IntArray, k: Int): IntArray {
        // Считаем кол-во вхождений элементов
        val freqs: MutableMap<Int, Int> = HashMap(nums.size)
        for (num in nums) {
            freqs[num] = freqs.getOrDefault(num, 0) + 1
        }

        // Формируем очередь с приоритетом
        val pq = PriorityQueue<Map.Entry<Int, Int>>(compareBy { it.value })
        for (pair in freqs.entries) {
            pq.add(pair)

            // Храним в очереди только k элементов, чтобы итоговая вычислительная сложность была O(n*log(k))
            // Редкие элементы в итоге будут вытеснены первыми
            if (pq.size > k) {
                pq.remove()
            }
        }

        // Забираем оставшиеся k элементов из очереди с приоритетом в обратном порядке
        val result = IntArray(k)
        for (i in k - 1 downTo 0) {
            val top = pq.remove()
            result[i] = top.key
        }

        return result
    }
}

val res = Solution().topKFrequent(intArrayOf(1, 1, 1, 2, 2, 3), 2)
println(res.contentToString())
