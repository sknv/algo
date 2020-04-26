const swap = (arr, i, j) => {
    const tmp = arr[i]
    arr[i] = arr[j]
    arr[j] = tmp
}

/**
 * @param {number[]} nums
 * @return {number[]}
 */
const bubbleSort = (nums) => {
    let swapped = false
    for (let i = 0; i < nums.length; i++) {
        swapped = false
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
    return nums
}
