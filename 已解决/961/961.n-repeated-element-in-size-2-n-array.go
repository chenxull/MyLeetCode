package problem961

/*
 * @lc app=leetcode id=961 lang=golang
 *
 * [961] N-Repeated Element in Size 2N Array
 *
 * https://leetcode.com/problems/n-repeated-element-in-size-2n-array/description/
 *
 * algorithms
 * Easy (73.46%)
 * Total Accepted:    43K
 * Total Submissions: 58.9K
 * Testcase Example:  '[1,2,3,3]'
 *
 * In a array A of size 2N, there are N+1 unique elements, and exactly one of
 * these elements is repeated N times.
 *
 * Return the element repeated N times.
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
 * Input: [1,2,3,3]
 * Output: 3
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [2,1,2,5,3,2]
 * Output: 2
 *
 *
 *
 * Example 3:
 *
 *
 * Input: [5,1,5,2,5,3,5,4]
 * Output: 5
 *
 *
 *
 *
 * Note:
 *
 *
 * 4 <= A.length <= 10000
 * 0 <= A[i] < 10000
 * A.length is even
 *
 *
 *
 *
 *
 */
func repeatedNTimes(A []int) int {
	hasseen := [10000]bool{}
	var val int
	for _, val = range A {
		if hasseen[val] {
			break
		}
		hasseen[val] = true

	}
	return val
}
