package main

import "fmt"

/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 *
 * https://leetcode.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (53.51%)
 * Total Accepted:    434.1K
 * Total Submissions: 806.8K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * Given an array nums, write a function to move all 0's to the end of it while
 * maintaining the relative order of the non-zero elements.
 *
 * Example:
 *
 *
 * Input: [0,1,0,3,12]
 * Output: [1,3,12,0,0]
 *
 * Note:
 *
 *
 * You must do this in-place without making a copy of the array.
 * Minimize the total number of operations.
 *
 */
func moveZeroes(nums []int) []int {
	if nums == nil {
		return nil
	}
	lastNonZeroFoundAt := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZeroFoundAt] = nums[i]
			lastNonZeroFoundAt++
		}
	}

	for i := lastNonZeroFoundAt; i < len(nums); i++ {
		nums[i] = 0
	}
	return nums
}

func main() {
	array := []int{0, 0, 0}
	fmt.Println(moveZeroes(array))
}
