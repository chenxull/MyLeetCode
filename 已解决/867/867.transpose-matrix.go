package main

import "fmt"

/*
 * @lc app=leetcode id=867 lang=golang
 *
 * [867] Transpose Matrix
 *
 * https://leetcode.com/problems/transpose-matrix/description/
 *
 * algorithms
 * Easy (63.64%)
 * Total Accepted:    39.3K
 * Total Submissions: 61.7K
 * Testcase Example:  '[[1,2,3],[4,5,6],[7,8,9]]'
 *
 * Given a matrix A, return the transpose of A.
 *
 * The transpose of a matrix is the matrix flipped over it's main diagonal,
 * switching the row and column indices of the matrix.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [[1,2,3],[4,5,6],[7,8,9]]
 * Output: [[1,4,7],[2,5,8],[3,6,9]]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,2,3],[4,5,6]]
 * Output: [[1,4],[2,5],[3,6]]
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 1000
 * 1 <= A[0].length <= 1000
 *
 *
 *
 *
 */

/*
	在 go 语言中如何定义一个二维数组。
	矩阵转置之后行和列要发生变化，这点不考虑到会导致数组越界
*/
func transpose(A [][]int) [][]int {

	row := len(A)
	col := len(A[0])
	ans := make([][]int, col)
	for i := 0; i < col; i++ {
		ans[i] = make([]int, row)
	}
	fmt.Println("row:", len(ans), "col:", len(ans[0]))
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			ans[j][i] = A[i][j]
		}
	}
	return ans
}

func main() {
	array := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		//{7, 8, 9},
	}
	fmt.Println(transpose(array))
}
