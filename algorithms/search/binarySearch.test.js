import binarySearch from './binarySearch.js'

const testBinarySearch = () => {
    const nums = [5, -2, 12, 0, 9]

    let idx = binarySearch(nums, 12), expected = 2
    console.assert(idx === expected, `expect to be: ${expected}, got instead: ${idx}`)

    idx = binarySearch(nums, 6), expected = -1
    console.assert(idx === expected, `expect to be: ${expected}, got instead: ${idx}`)
}

testBinarySearch()
