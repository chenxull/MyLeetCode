package main

import "fmt"

/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 *
 * https://leetcode.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (41.84%)
 * Total Accepted:    189.4K
 * Total Submissions: 446.9K
 * Testcase Example:  '3'
 *
 * Given a non-negative index k where k ≤ 33, return the k^th index row of the
 * Pascal's triangle.
 *
 * Note that the row index starts from 0.
 *
 *
 * In Pascal's triangle, each number is the sum of the two numbers directly
 * above it.
 *
 * Example:
 *
 *
 * Input: 3
 * Output: [1,3,3,1]
 *
 *
 * Follow up:
 *
 * Could you optimize your algorithm to use only O(k) extra space?
 *
 */
func getRow(rowIndex int) []int {
	var array []int = make([]int, rowIndex+1)
	array[0] = 1
	for i := 0; i < rowIndex+1; i++ {
		for j := i; j > 0; j-- {
			array[j] += array[j-1]
		}
	}
	return array
}

func main() {
	array := getRow(3)
	fmt.Println(array)
}
