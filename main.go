package main

import "fmt"

func pivotIndex(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	left := 0
	for k, v := range nums {
		if 2*left+v == totalSum {
			return k
		}
		left += v
	}
	return -1
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
