package main

import "fmt"

/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 *
 * https://leetcode.com/problems/contains-duplicate/description/
 *
 * algorithms
 * Easy (50.65%)
 * Total Accepted:    311.9K
 * Total Submissions: 609.9K
 * Testcase Example:  '[1,2,3,1]'
 *
 * Given an array of integers, find if the array contains any duplicates.
 *
 * Your function should return true if any value appears at least twice in the
 * array, and it should return false if every element is distinct.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,1]
 * Output: true
 *
 * Example 2:
 *
 *
 * Input: [1,2,3,4]
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: [1,1,1,3,3,4,3,2,4,2]
 * Output: true
 *
 */
func containsDuplicate(nums []int) bool {
	mp := make(map[int][]int)
	for key, val := range nums {
		mp[val] = append(mp[val], key)
		if len(mp[val]) > 1 {
			return true
		}
	}
	return false
}

func main() {
	array := []int{1, 2, 3, 4}
	if containsDuplicate(array) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
