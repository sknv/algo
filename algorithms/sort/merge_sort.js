const swap = (arr, i, j) => {
    const tmp = arr[i]
    arr[i] = arr[j]
    arr[j] = tmp
}

/**
 * @param {number[]} arr
 * @param {number} left
 * @param {number} mid
 * @param {number} right
 */
const merge = (arr, left, mid, right) => {
    // Create temp slices
    const leftSliceEnd = mid + 1, rightSliceEnd = right + 1
    const leftSlice = arr.slice(left, leftSliceEnd), rightSlice = arr.slice(leftSliceEnd, rightSliceEnd)

    // Merge the temp slices
    let i = 0, j = 0, mergedIndex = left
    while (i < leftSlice.length && j < rightSlice.length) {
        if (leftSlice[i] < rightSlice[j]) {
            arr[mergedIndex] = leftSlice[i]
            i++
        } else {
            arr[mergedIndex] = rightSlice[j]
            j++
        }
        mergedIndex++
    }

    // Copy remaining elements
    while (i < leftSlice.length) {
        arr[mergedIndex] = leftSlice[i]
        i++
        mergedIndex++
    }
    while (j < rightSlice.length) {
        arr[mergedIndex] = rightSlice[j]
        j++
        mergedIndex++
    }
}

/**
 * @param {number[]} arr
 * @param {number} left
 * @param {number} right
 */
const mergeSort = (arr, left, right) => {
    if (left < right) {
        const mid = Math.trunc((left + right) / 2)

        mergeSort(arr, left, mid)
        mergeSort(arr, mid + 1, right)
        merge(arr, left, mid, right)
    }
}

/**
 * @param {number[]} nums
 * @return {number[]}
 */
const sortArray = nums => {
    mergeSort(nums, 0, nums.length - 1)
    return nums
}
