package main

import (
	"LeetCode/sort"
	"fmt"
	"time"
)

func shell(nums []int) ([]int, time.Duration) {
	if len(nums) == 0 {
		return nil, 0
	}
	t1 := time.Now()
	N := len(nums)
	h := 1
	// 获取合适的步长
	for h < N/3 {
		h = h*3 + 1
	}

	for h >= 1 {
		// 以 h 为起点，遍历处理整个数组
		for i := h; i < N; i++ {
			// 在排序的过程中，以 h 为步长，向后处理
			for j := i; j >= h; j -= h {
				if nums[j-h] > nums[j] {
					utils.Swap(nums, j, j-h)
				}
			}
		}
		h = h / 3
	}

	during := time.Since(t1)
	return nums, during
}

func main() {
	nums := utils.ArrayRand(100000)
	fmt.Println(shell(nums))
}
