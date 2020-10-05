package search

func BinarySearchInt(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		cur := arr[mid]

		if cur == target {
			return mid
		}

		if target < cur {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
