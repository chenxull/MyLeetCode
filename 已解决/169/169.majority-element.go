package main

import "fmt"

/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 *
 * https://leetcode.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (51.35%)
 * Total Accepted:    357.5K
 * Total Submissions: 690.4K
 * Testcase Example:  '[3,2,3]'
 *
 * Given an array of size n, find the majority element. The majority element is
 * the element that appears more than ⌊ n/2 ⌋ times.
 *
 * You may assume that the array is non-empty and the majority element always
 * exist in the array.
 *
 * Example 1:
 *
 *
 * Input: [3,2,3]
 * Output: 3
 *
 * Example 2:
 *
 *
 * Input: [2,2,1,1,1,2,2]
 * Output: 2
 *
 *
 */
func majorityElement(nums []int) int {

	count := 0
	var candidate int
	for _, val := range nums {
		if count == 0 {
			candidate = val
		}
		if val == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

func main() {
	array := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement(array))
}
