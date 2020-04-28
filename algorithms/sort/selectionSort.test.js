import selectionSort from './selectionSort.js'

const testSelectionSort = () => {
    let nums = [5, -2, 12, 0, 9], expected = [-2, 0, 5, 9, 12]
    selectionSort(nums)
    console.assert(JSON.stringify(nums) === JSON.stringify(expected), `expect to be: ${expected}, got instead: ${nums}`)

    nums = [1, -2, 10, -6], expected = [-6, -2, 1, 10]
    selectionSort(nums)
    console.assert(JSON.stringify(nums) === JSON.stringify(expected), `expect to be: ${expected}, got instead: ${nums}`)
}

testSelectionSort()
