package sort

// MergeSort 归并排序
func MergeSort(list []int) {
	if len(list) == 1 {
		return
	}

	mid := len(list) / 2
	MergeSort(list[:mid])
	MergeSort(list[mid:])

	merge(list)
}

// merge 合并数组
func merge(list []int) {
	temp := make([]int, len(list))
	for i, v := range list {
		temp[i] = v
	}

	mid := len(list) / 2
	l, r := 0, mid

	var index int
	for l < mid && r < len(list) {
		if temp[l] < temp[r] {
			list[index] = temp[l]
			l++
		} else {
			list[index] = temp[r]
			r++
		}

		index++
	}

	for i, v := range temp[l:mid] {
		list[index+i] = v
	}
	for i, v := range temp[r:] {
		list[index+i] = v
	}

	return
}
