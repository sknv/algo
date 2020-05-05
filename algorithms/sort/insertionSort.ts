function swap(arr: number[], i: number, j: number) {
    [arr[i], arr[j]] = [arr[j], arr[i]]
}

function insertionSort(arr: number[]) {
    for (let i = 0; i < arr.length; i++) {
        for (let j = i; j > 0; j--) {
            if (arr[j] < arr[j - 1]) {
                swap(arr, j, j - 1)
            } else {
                break
            }
        }
    }
}

export default insertionSort
