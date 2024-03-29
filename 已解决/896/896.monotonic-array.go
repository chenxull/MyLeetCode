package main

import "fmt"

/*
 * @lc app=leetcode id=896 lang=golang
 *
 * [896] Monotonic Array
 *
 * https://leetcode.com/problems/monotonic-array/description/
 *
 * algorithms
 * Easy (54.51%)
 * Total Accepted:    35.1K
 * Total Submissions: 64.1K
 * Testcase Example:  '[1,2,2,3]'
 *
 * An array is monotonic if it is either monotone increasing or monotone
 * decreasing.
 *
 * An array A is monotone increasing if for all i <= j, A[i] <= A[j].  An array
 * A is monotone decreasing if for all i <= j, A[i] >= A[j].
 *
 * Return true if and only if the given array A is monotonic.
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1,2,2,3]
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [6,5,4,4]
 * Output: true
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [1,3,2]
 * Output: false
 *
 *
 *
 * Example 4:
 *
 *
 * Input: [1,2,4,5]
 * Output: true
 *
 *
 *
 * Example 5:
 *
 *
 * Input: [1,1,1]
 * Output: true
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 50000
 * -100000 <= A[i] <= 100000
 *
 *
 *
 *
 *
 *
 *
 */
func isMonotonic(A []int) bool {
	if A == nil {
		return false
	}
	increasing := true
	decreasing := true

	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			increasing = false
		}
		if A[i] < A[i+1] {
			decreasing = false
		}
	}

	return increasing || decreasing
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	if isMonotonic(array) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
