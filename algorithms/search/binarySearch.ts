/**
 * Return the index of an element if exists, -1 otherwise
 */
function binarySearch(arr: number[], target: number): number {
    let left = 0, right = arr.length - 1
    while (left <= right) {
        const mid = left + Math.trunc((right - left) / 2)

        if (target === arr[mid]) {
            return mid
        }

        if (target < arr[mid]) {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return -1
}

export default binarySearch
