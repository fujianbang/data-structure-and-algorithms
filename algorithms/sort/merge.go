package sort

// MergeSort 归并排序
func MergeSort(list []int) {
	mergeSortFunc(list, 0, len(list)-1)
}

// mergeSortFunc 归并排序子算法函数
func mergeSortFunc(list []int, p, q int) {
	if p >= q {
		return
	}

	mid := (p + q) / 2
	mergeSortFunc(list, p, mid)
	mergeSortFunc(list, mid+1, q)

	merge(list, p, q)
}

// merge 合并数组
func merge(list []int, left, right int) {

	temp := make([]int, right-left+1)

	for i := left; i <= right; i++ {
		temp[i-left] = list[i]
	}

	m := (right-left)/2 + 1
	l, r := 0, m

	idx := left
	for l < m && r < len(list) {
		if temp[l] < temp[r] {
			list[idx] = temp[l]
			l++
		} else {
			list[idx] = temp[r]
			r++
		}

		idx++
	}

	for i, v := range temp[left:m] {
		list[idx+i] = v
	}
	for i, v := range temp[r:] {
		list[idx+i] = v
	}

	return
}
