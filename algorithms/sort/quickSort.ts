function swap(arr: number[], i: number, j: number) {
    [arr[i], arr[j]] = [arr[j], arr[i]]
}

/**
 * Returns partitioning index.
 */
function partition(arr: number[], left: number, right: number): number {
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

function quickSort(arr: number[], left: number, right: number) {
    if (left < right) {
        let pi = partition(arr, left, right) // partitioning index

        quickSort(arr, left, pi - 1)
        quickSort(arr, pi + 1, right)
    }
}

function sortArray(arr: number[]) {
    quickSort(arr, 0, arr.length - 1)
}

export default sortArray
