package main

import "fmt"

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 *
 * https://leetcode.com/problems/two-sum/description/
 *
 * algorithms
 * Easy (40.80%)
 * Total Accepted:    1.4M
 * Total Submissions: 3.5M
 * Testcase Example:  '[2,7,11,15]\n9'
 *
 * Given an array of integers, return indices of the two numbers such that they
 * add up to a specific target.
 *
 * You may assume that each input would have exactly one solution, and you may
 * not use the same element twice.
 *
 * Example:
 *
 *
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1].
 *
 *
 *
 *
 */
func twoSum(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return nums
	}
	mp := make(map[int]int)

	for i, num := range nums {
		fmt.Println(i, num)
		if _, ok := mp[target-num]; ok {
			return []int{mp[target-num], i}
		}
		mp[num] = i
	}
	return []int{}
}

func main() {
	nums := []int{3, 3}
	fmt.Println(twoSum(nums, 6))

}
