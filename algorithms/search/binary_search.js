/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number} Index of element if exists, -1 otherwise
 */
const binarySearch = (nums, target) => {
    let left = 0, right = nums.length - 1
    while (left <= right) {
        const mid = left + Math.trunc((right - left) / 2)
        const check = nums[mid]

        if (check === target) {
            return mid
        }
        if (target < check) {
            right = mid - 1
            continue
        }
        left = mid + 1
    }
    return -1
}
