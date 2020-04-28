const swap = (arr, i, j) => {
    const tmp = arr[i]
    arr[i] = arr[j]
    arr[j] = tmp
}

/**
 * @param {number[]} arr
 */
const bubbleSort = arr => {
    for (let i = 0; i < arr.length; i++) {
        let swapped = false
        for (let j = 0; j < arr.length - i - 1; j++) {
            if (arr[j] > arr[j + 1]) {
                swap(arr, j, j + 1)
                swapped = true
            }
        }
        if (!swapped) {
            break
        }
    }
}

export default bubbleSort
