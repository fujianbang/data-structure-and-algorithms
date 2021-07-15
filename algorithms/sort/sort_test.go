package sort

import (
	"sort"
	"testing"
)

func getTestCases() []sort.IntSlice {
	cases := []sort.IntSlice{
		{2, 5, 6, 7, 4, 8, 9, 10},
		{1, 2, 3, 3, 4, 5},
		{3, 4, 1, 5, 0, 1, 2},
		{9, 11, 17, 2, 6, 0, 10, 3, 13, 16, 18, 15, 5, 14, 4, 12, 8, 19, 1, 7},
		{6, 7, 4, 13, 0, 8, 1, 3, 11, 16, 9, 19, 18, 15, 2, 12, 17, 14, 10, 5},
		{7, 9, 3, 10, 16, 13, 17, 1, 11, 8, 5, 4, 12, 18, 6, 0, 19, 14, 2, 15},
		{19, 18, 6, 3, 10, 1, 5, 4, 12, 7, 2, 13, 14, 15, 9, 8, 0, 16, 11, 17},
		{17, 16, 3, 6, 8, 4, 14, 11, 18, 15, 10, 7, 12, 9, 0, 5, 19, 13, 2, 1},
		{11, 9, 1, 5, 7, 8, 3, 18, 2, 13, 14, 16, 17, 15, 19, 4, 6, 10, 12, 0},
		{12, 18, 8, 4, 19, 15, 14, 5, 0, 16, 9, 2, 11, 3, 6, 17, 10, 7, 13, 1},
		{5, 8, 11, 18, 1, 13, 0, 6, 12, 10, 4, 14, 16, 3, 15, 19, 2, 9, 17, 7},
		{18, 3, 0, 8, 12, 7, 5, 10, 19, 11, 2, 1, 13, 6, 17, 9, 14, 4, 15, 16},
	}

	return cases
}

// TestBubbleSort 测试冒泡排序
func TestBubbleSort(t *testing.T) {
	for _, tt := range getTestCases() {
		item := make([]int, len(tt))
		copy(item, tt)
		BubbleSort(tt)

		if !sort.IsSorted(tt) {
			t.Errorf("bubble(%v) = %v", item, tt)
		}
	}
}

// TestMergeSort 测试归并排序
func TestMergeSort(t *testing.T) {
	for _, tt := range getTestCases() {
		item := make([]int, len(tt))
		copy(item, tt)
		MergeSort(tt)

		if !sort.IsSorted(tt) {
			t.Errorf("merge(%v) = %v", item, tt)
		}
	}
}

func TestQuickSort(t *testing.T) {
	for _, tt := range getTestCases() {
		item := make([]int, len(tt))
		copy(item, tt)
		QuickSort(tt)

		if !sort.IsSorted(tt) {
			t.Errorf("merge(%v) = %v", item, tt)
		}
	}
}
