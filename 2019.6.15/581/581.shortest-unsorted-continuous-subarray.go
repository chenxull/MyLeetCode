package problem581

import "sort"

/*
 * @lc app=leetcode id=581 lang=golang
 *
 * [581] Shortest Unsorted Continuous Subarray
 */
// 和快拍的思路相似，维护二个标识 分别指向最大值和最小值。分别从左右遍历一遍数组。
func findUnsortedSubarray(nums []int) int {

	numSorted := make([]int, len(nums))
	copy(numSorted, nums)
	sort.Ints(numSorted)

	start, end := len(nums), 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != numSorted[i] {
			start = min(start, i)
			end = max(end, i)
		}
	}

	if end > start {
		return end - start + 1
	}

	return 0

	// n := len(nums)
	// left, right := 0, -1 // ?
	// max, min := nums[0], nums[n-1]
	// for i := 1; i < n; i++ {
	// 	if max <= nums[i] {
	// 		max = nums[i]
	// 	} else {
	// 		right = i
	// 	}

	// 	j := n - 1 - i
	// 	if min >= nums[j] {
	// 		min = nums[j]
	// 	} else {
	// 		left = j
	// 	}
	// }
	// return right - left + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
