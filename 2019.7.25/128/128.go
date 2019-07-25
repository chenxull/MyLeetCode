package problem128

import "sort"

/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */
//   构建一个 map 来存储出现的元素
func longestConsecutive1(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	sort.Ints(nums)
	temp, max := 1, 0

	for i := 1; i < len(nums); i++ {
		if nums[i]-1 == nums[i-1] {
			temp++

		} else if nums[i] != nums[i-1] {
			// 当遇到二个数字不同时，才将 temp 置为 1 ，如果二个数字相同，应该判断一下最大值
			temp = 1
		}
		if max < temp {
			max = temp
		}
	}
	return max

}
