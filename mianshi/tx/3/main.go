package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	var length int
	fmt.Scan(&n, &length)

	nums := make([][]int, 0, n)
	for i := range nums {
		nums[i] = make([]int, 0, 2)
	}

	for i := 0; i < n; i++ {
		left, right := 0, 0
		fmt.Scan(&left, &right)
		temp := []int{left, right}
		nums = append(nums, temp)
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i][0] < nums[j][0]
	})

	if nums[0][0] != 0 {
		fmt.Println(-1)
	}

	temp := nums[0]
	count := 1
	for i := 1; i < n; i++ {
		if nums[i][0] <= temp[1] {
			continue
		}
		count++
		temp = nums[i]
		if nums[i][1] >= length {
			fmt.Println(count)
			return
		}
	}

	fmt.Println(count)
}
