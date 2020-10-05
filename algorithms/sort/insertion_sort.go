package sort

func InsertionSortInt(arr []int, less func(cur, next int) bool) {
	for i := 0; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if !less(arr[j-1], arr[j]) {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			} else {
				break
			}
		}
	}
}
