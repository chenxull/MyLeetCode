package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	nums := make([]int, n*m)
	index := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			nums[index] = (i + 1) * (j + 1)
			index++
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	// fmt.Println(nums)
	fmt.Println(nums[k-1])

}
