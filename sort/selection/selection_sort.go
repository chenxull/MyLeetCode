package main

import (
	"LeetCode/sort"
	"fmt"
	"time"
)

//  选择排序

func selecttionSort(nums []int) ([]int, time.Duration) {
	if len(nums) == 0 {
		return nil, 0
	}
	t1 := time.Now()
	n := len(nums)

	for i := 0; i < n-1; i++ {
		min := i // 用来存储最小的下标
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		utils.Swap(nums, i, min)
	}

	during := time.Since(t1)
	return nums, during
}

func main() {
	nums := utils.ArrayRand(100000)
	fmt.Println(selecttionSort(nums))
}
