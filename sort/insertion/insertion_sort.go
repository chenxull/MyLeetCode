package main

import (
	"LeetCode/sort"
	"fmt"
	"time"
)

func insertion(nums []int) ([]int, time.Duration) {
	if len(nums) == 0 {
		return nil, 0
	}
	t1 := time.Now()
	for i := 1; i < len(nums); i++ {
		// [0,j] 数组有序,
		for j := i; j > 0; j-- {
			if nums[j] < nums[j-1] {
				utils.Swap(nums, j, j-1)
			}
		}
	}
	during := time.Since(t1)
	return nums, during

}

func main() {
	nums := utils.ArrayRand(100000)
	fmt.Println(insertion(nums))
}
