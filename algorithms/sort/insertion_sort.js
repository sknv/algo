const swap = (arr, i, j) => {
    const tmp = arr[i]
    arr[i] = arr[j]
    arr[j] = tmp
}

/**
 * @param {number[]} nums
 */
const insertionSort = nums => {
    for (let i = 0; i < nums.length; i++) {
        for (let j = i; j > 0; j--) {
            if (nums[j] < nums[j - 1]) {
                swap(nums, j, j - 1)
                continue
            }
            break
        }
    }
}

module.exports = insertionSort
