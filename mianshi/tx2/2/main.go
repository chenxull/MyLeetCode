package main

import (
	"fmt"
	"sort"
)

/*
2
1 1
2 2
*/
func main() {
	var n int
	fmt.Scan(&n)

	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, 2)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			tmp := 0
			fmt.Scan(&tmp)
			nums[i][j] = tmp
		}
	}
	// fmt.Println(nums)
	// 使用二次排序求解
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] > nums[j][0]
	})
	sort.Slice(nums, func(i, j int) bool {
		return nums[i][1] < nums[j][1]
	})

	ans := 0
	for i := 0; i < n; i++ {
		ans += nums[i][0]*i + nums[i][1]*(n-i-1)
	}
	fmt.Println(ans)
	// fmt.Println(nums)
}
