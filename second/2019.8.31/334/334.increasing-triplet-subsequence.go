package problem334

import "math"

/*
 * @lc app=leetcode id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 */
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	num1 := nums[0]
	num2 := math.MaxInt64
	for i := 1; i < len(nums); i++ {
		if nums[i] < num1 {
			num1 = nums[i]
		} else {
			if num2 < nums[i] {
				return true
			}
			num2 = nums[i]

		}
	}
	return false
}
