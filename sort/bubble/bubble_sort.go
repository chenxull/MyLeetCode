package main

import (
	"LeetCode/sort"
	"fmt"
	"time"
)

func bubble(nums []int) ([]int, time.Duration) {
	if len(nums) == 0 {
		return nil, 0
	}
	t1 := time.Now()

	n := len(nums)
	isSorted := false
	for i := n - 1; i > 0 && !isSorted; i-- {
		isSorted = true
		for j := 0; j < i; j++ {
			// 当前元素大于后面的元素,冒泡交换
			if nums[j+1] < nums[j] {
				isSorted = false
				utils.Swap(nums, j, j+1)
			}
		}
	}

	during := time.Since(t1)
	return nums, during
}

func main() {
	nums := utils.ArrayRand(100000)
	fmt.Println(bubble(nums))
}
