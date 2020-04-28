/**
 * @param {number[]} arr
 * @param {number} target
 * @returns {number} Index of element if exists, -1 otherwise
 */
const binarySearch = (arr, target) => {
    let left = 0, right = arr.length - 1
    while (left <= right) {
        const mid = left + Math.trunc((right - left) / 2)

        if (arr[mid] === target) {
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
