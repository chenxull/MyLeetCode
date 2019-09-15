package main

import "fmt"

func minNumber(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	l, h := 0, len(nums)-1
	for l < h {
		m := l + (h-l)/2

		if nums[m] == target {
			return m
		} else if nums[m] <= nums[h] {
			h = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	nums := []int{2, 5, 6, 3, 9, 1}
	fmt.Println(minNumber(nums, 9))

}
