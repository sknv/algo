/**
 * @param {number[]} arr
 * @param {number[]} leftSlice
 * @param {number[]} rightSLice
 */
const merge = (arr, leftSlice, rightSlice) => {
    // Merge the slices
    let i = 0, j = 0, mergeIndex = 0
    while (i < leftSlice.length && j < rightSlice.length) {
        if (leftSlice[i] < rightSlice[j]) {
            arr[mergeIndex] = leftSlice[i]
            i++
        } else {
            arr[mergeIndex] = rightSlice[j]
            j++
        }
        mergeIndex++
    }

    // Copy remaining elements
    while (i < leftSlice.length) {
        arr[mergeIndex] = leftSlice[i]
        i++
        mergeIndex++
    }
    while (j < rightSlice.length) {
        arr[mergeIndex] = rightSlice[j]
        j++
        mergeIndex++
    }
}

/**
 * @param {number[]} arr
 */
const mergeSort = arr => {
    if (arr.length <= 1) {
        return
    }

    // Split array on two halves
    const mid = Math.trunc(arr.length / 2)
    const leftSlice = arr.slice(0, mid), rightSlice = arr.slice(mid, arr.length) // get copies of array halves

    // Sort two halves of splitted array
    mergeSort(leftSlice)
    mergeSort(rightSlice)

    // Merge two sorted slices into one
    merge(arr, leftSlice, rightSlice)
}

module.exports = mergeSort
