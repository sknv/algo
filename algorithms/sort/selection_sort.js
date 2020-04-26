const swap = (arr, i, j) => {
    const tmp = arr[i]
    arr[i] = arr[j]
    arr[j] = tmp
}

/**
 * @param {number[]} nums
 */
const selectionSort = nums => {
    for (let i = 0; i < nums.length - 1; i++) {
        let minIndex = i
        for (let j = i + 1; j < nums.length; j++) {
            if (nums[j] < nums[minIndex]) {
                minIndex = j
            }
        }
        swap(nums, i, minIndex)
    }
}

module.exports = selectionSort
