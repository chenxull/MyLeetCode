package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=561 lang=golang
 *
 * [561] Array Partition I
 *
 * https://leetcode.com/problems/array-partition-i/description/
 *
 * algorithms
 * Easy (68.23%)
 * Total Accepted:    133.1K
 * Total Submissions: 194.5K
 * Testcase Example:  '[1,4,3,2]'
 *
 *
 * Given an array of 2n integers, your task is to group these integers into n
 * pairs of integer, say (a1, b1), (a2, b2), ..., (an, bn) which makes sum of
 * min(ai, bi) for all i from 1 to n as large as possible.
 *
 *
 * Example 1:
 *
 * Input: [1,4,3,2]
 *
 * Output: 4
 * Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3,
 * 4).
 *
 *
 *
 * Note:
 *
 * n is a positive integer, which is in the range of [1, 10000].
 * All the integers in the array will be in the range of [-10000, 10000].
 *
 *
 */
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var MaxSum int
	for i := 0; i < len(nums); i += 2 {
		MaxSum += nums[i]
	}
	return MaxSum

}

func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arrayPairSum(array))
}
