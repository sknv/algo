package sort

func BubbleSortInt(arr []int, less func(cur, next int) bool) {
	for i := 0; i < len(arr)-1; i++ {
		swapped := false

		for j := 0; j < len(arr)-i-1; j++ {
			if !less(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}

		if !swapped {
			return
		}
	}
}
