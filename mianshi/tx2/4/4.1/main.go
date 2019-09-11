package main

import (
	"fmt"
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
	// dp := make([]int, n)
	M := 0
	sum := make([]int, n)
	for i := 0; i < n; i++ {
		sum[i] = findSum(nums, i)
		M = max(M, sum[i]*nums[i])
		// fmt.Println(M)
	}
	fmt.Println(M)

}

func findSum(nums []int, i int) int {
	min := nums[i]
	sum := nums[i]
	if i == 0 && nums[0] > nums[1] {
		return nums[0]
	}
	j := i
	for i < len(nums)-1 && min < nums[i+1] {
		sum += nums[i+1]
		i++
	}

	for j > 0 && min < nums[j-1] {
		sum += nums[j-1]
		j--
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
