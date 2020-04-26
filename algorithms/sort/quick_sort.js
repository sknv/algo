const swap = (arr, i, j) => {
    const tmp = arr[i]
    arr[i] = arr[j]
    arr[j] = tmp
}

/**
 * @param {number[]} arr
 * @param {number} left
 * @param {number} right
 * @return {number} Partitioning index
 */
const partition = (arr, left, right) => {
    const pivot = arr[right] // take the right element as the pivot one
    let partIndex = left
    for (let i = left; i < right; i++) { // place all the elements smaller than pivot to the left
        if (arr[i] < pivot) {
            swap(arr, partIndex, i)
            partIndex++
        }
    }
    swap(arr, partIndex, right) // place all the elements greater than pivot to the right
    return partIndex
}

/**
 * @param {number[]} arr
 * @param {number} left
 * @param {number} right
 */
const quickSort = (arr, left, right) => {
    if (left < right) {
        let pi = partition(arr, left, right) // partitioning index

        quickSort(arr, left, pi - 1)
        quickSort(arr, pi + 1, right)
    }
}

/**
 * @param {number[]} nums
 * @return {number[]}
 */
const sortArray = nums => {
    quickSort(nums, 0, nums.length - 1)
    return nums
}
