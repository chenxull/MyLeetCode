package main

import "fmt"

func duplicate(nums []int, length int) int {
	if len(nums) == 0 {
		return -1
	}
	res := 0
	for i := 0; i < length; i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				res = nums[i]
				return res
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}

	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5}
	fmt.Println(duplicate(nums, len(nums)))
}
