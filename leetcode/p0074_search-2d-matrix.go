package leetcode

// https://leetcode.com/problems/search-a-2d-matrix/description/

func searchMatrix(matrix [][]int, target int) bool {
	firstElements := make([]int, len(matrix))
	for i, row := range matrix {
		firstElements[i] = row[0]
	}

	rowToSearch, exact := binarySearchInMatrixRow(firstElements, target)
	if exact {
		return true
	}

	_, found := binarySearchInMatrixRow(matrix[rowToSearch], target)

	return found
}

func binarySearchInMatrixRow(array []int, target int) (int, bool) {
	left, right := 0, len(array)-1
	for left <= right {
		mid := left + (right-left)/2
		value := array[mid]

		if value == target {
			return mid, true
		}

		if target < value {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return max(right, 0), false
}
