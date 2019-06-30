package main

import "fmt"

/*
 * @lc app=leetcode id=922 lang=golang
 *
 * [922] Sort Array By Parity II
 *
 * https://leetcode.com/problems/sort-array-by-parity-ii/description/
 *
 * algorithms
 * Easy (66.65%)
 * Total Accepted:    32.1K
 * Total Submissions: 48.1K
 * Testcase Example:  '[4,2,5,7]'
 *
 * Given an array A of non-negative integers, half of the integers in A are
 * odd, and half of the integers are even.
 *
 * Sort the array so that whenever A[i] is odd, i is odd; and whenever A[i] is
 * even, i is even.
 *
 * You may return any answer array that satisfies this condition.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [4,2,5,7]
 * Output: [4,5,2,7]
 * Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been
 * accepted.
 *
 *
 *
 *
 * Note:
 *
 *
 * 2 <= A.length <= 20000
 * A.length % 2 == 0
 * 0 <= A[i] <= 1000
 *
 *
 *
 *
 *
 */
func sortArrayByParityII(A []int) []int {
	len := len(A)
	ans := make([]int, len)
	oddCount := 1
	evenCount := 0
	for _, val := range A {
		if val%2 == 1 {
			ans[oddCount] = val
			oddCount += 2
		} else {
			ans[evenCount] = val
			evenCount += 2
		}
	}
	return ans
}
func main() {
	array := []int{4, 2, 5, 7}
	fmt.Println(sortArrayByParityII(array))
}
