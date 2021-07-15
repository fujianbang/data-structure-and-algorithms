package sort

import (
	"testing"
)

func TestPartition(t *testing.T) {
	l := []int{6, 11, 3, 9, 8}

	n := partition(l, 0, len(l)-1)

	t.Log(n)
	t.Log(l)
}
