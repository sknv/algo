/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number} Index of element if exists, -1 otherwise
 */
const binarySearch = (nums, target) => {
    let left = 0, right = nums.length - 1
    while (left <= right) {
        const mid = left + Math.trunc((right - left) / 2)

        if (nums[mid] === target) {
            return mid
        }
        if (target < nums[mid]) {
            right = mid - 1
            continue
        }
        left = mid + 1
    }
    return -1
}

module.exports = binarySearch
