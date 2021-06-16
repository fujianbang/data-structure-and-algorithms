package sort

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestMerge(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{[]int{2, 5, 6, 7, 4, 8, 9, 10}, []int{2, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 4, 5, 0, 1, 2}, []int{0, 1, 2, 3, 4, 5}},
	}

	for _, tt := range cases {
		item := make([]int, len(tt.input))
		copy(item, tt.input)
		merge(item)
		if !equal(item, tt.expected) {
			t.Errorf("merge(%v) = %v; expected %v", tt.input, item, tt.expected)
		}
	}
}

func TestMergeSort(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{[]int{2, 5, 6, 7, 4, 8, 9, 10}, []int{2, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 4, 5, 0, 1, 2}, []int{0, 1, 2, 3, 4, 5}},
	}

	for _, tt := range cases {
		item := make([]int, len(tt.input))
		copy(item, tt.input)
		MergeSort(item)
		if !equal(item, tt.expected) {
			t.Errorf("merge(%v) = %v; expected %v", tt.input, item, tt.expected)
		}
	}
}

func TestMergeSort2(t *testing.T) {
	r := rand.Perm(20)
	t.Log(r)
	MergeSort(r)
	t.Log(r)
}
