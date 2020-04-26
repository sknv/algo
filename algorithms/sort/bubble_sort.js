const swap = (arr, i, j) => {
    const tmp = arr[i]
    arr[i] = arr[j]
    arr[j] = tmp
}

/**
 * @param {number[]} nums
 */
const bubbleSort = nums => {
    for (let i = 0; i < nums.length; i++) {
        let swapped = false
        for (let j = 0; j < nums.length - i - 1; j++) {
            if (nums[j] > nums[j + 1]) {
                swap(nums, j, j + 1)
                swapped = true
            }
        }
        if (!swapped) {
            break
        }
    }
}

module.exports = bubbleSort
