package main

import "fmt"

/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 *
 * https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (52.61%)
 * Total Accepted:    140.2K
 * Total Submissions: 265.2K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some
 * elements appear twice and others appear once.
 *
 * Find all the elements of [1, n] inclusive that do not appear in this array.
 *
 * Could you do it without extra space and in O(n) runtime? You may assume the
 * returned list does not count as extra space.
 *
 * Example:
 *
 * Input:
 * [4,3,2,7,8,2,3,1]
 *
 * Output:
 * [5,6]
 *
 *
 */
func findDisappearedNumbers(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	for i := 0; i < len(nums); i++ {
		val := abs(nums[i]) - 1
		if nums[val] > 0 {
			nums[val] = -nums[val]
		}
	}
	var ans []int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	array := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println("ans:", findDisappearedNumbers(array))
}
