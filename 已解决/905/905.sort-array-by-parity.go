package main

import "fmt"

/*
 * @lc app=leetcode id=905 lang=golang
 *
 * [905] Sort Array By Parity
 *
 * https://leetcode.com/problems/sort-array-by-parity/description/
 *
 * algorithms
 * Easy (72.09%)
 * Total Accepted:    74.1K
 * Total Submissions: 102.6K
 * Testcase Example:  '[3,1,2,4]'
 *
 * Given an array A of non-negative integers, return an array consisting of all
 * the even elements of A, followed by all the odd elements of A.
 *
 * You may return any answer array that satisfies this condition.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [3,1,2,4]
 * Output: [2,4,3,1]
 * The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 5000
 * 0 <= A[i] <= 5000
 *
 *
 *
 */

func sortArrayByParity(A []int) []int {
	leftPoint := 0
	rightPoint := len(A) - 1
	fmt.Println("DEBUG")
	for leftPoint < rightPoint {
		fmt.Println("debug")
		fmt.Println("leftpoint:", leftPoint)
		fmt.Println("rightPoint:", rightPoint)
		if A[leftPoint]%2 > A[rightPoint]%2 {
			temp := A[leftPoint]
			A[leftPoint] = A[rightPoint]
			A[rightPoint] = temp
		}
		if A[leftPoint]%2 == 0 {
			leftPoint++
		}
		if A[rightPoint]%2 == 1 {
			rightPoint--
		}
	}
	return A
}
func main() {
	array := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(array))
}
