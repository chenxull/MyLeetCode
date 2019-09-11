package main

import (
	"fmt"
	"math"
)

/*
5
7 2 4 6 5
*/
func main() {
	var n int
	fmt.Scan(&n)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		nums[i] = tmp
	}

	// sum := 0
	M := 0
	for i := 0; i < n; i++ {
		sum := nums[i]
		for j := i + 1; j <= n-1; j++ {
			min := findmin(nums, i, j)
			sum += nums[j]
			M = max(M, sum*min)
		}
	}
	fmt.Println(M)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findmin(nums []int, start, end int) int {
	min := math.MaxInt64

	for i := start; i <= end; i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
