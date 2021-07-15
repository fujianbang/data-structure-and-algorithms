package sort

// QuickSort 快速排序
func QuickSort(list []int) {
	quickSortFunc(list, 0, len(list)-1)
}

func quickSortFunc(list []int, left, right int) {
	if left >= right {
		return
	}
	n := partition(list, left, right)
	quickSortFunc(list, left, n-1)
	quickSortFunc(list, n+1, right)
}

func partition(list []int, left, right int) int {
	pivot := list[right]
	i := left
	for j := left; j <= right; j++ {
		if list[j] < pivot {
			if i != j {
				list[j], list[i] = list[i], list[j]
			}
			i++
		}
	}
	list[i], list[right] = list[right], list[i]
	return i
}
