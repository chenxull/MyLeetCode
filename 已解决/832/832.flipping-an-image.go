package main

import "fmt"

/*
 * @lc app=leetcode id=832 lang=golang
 *
 * [832] Flipping an Image
 *
 * https://leetcode.com/problems/flipping-an-image/description/
 *
 * algorithms
 * Easy (71.32%)
 * Total Accepted:    80K
 * Total Submissions: 111.5K
 * Testcase Example:  '[[1,1,0],[1,0,1],[0,0,0]]'
 *
 * Given a binary matrix A, we want to flip the image horizontally, then invert
 * it, and return the resulting image.
 *
 * To flip an image horizontally means that each row of the image is reversed.
 * For example, flipping [1, 1, 0] horizontally results in [0, 1, 1].
 *
 * To invert an image means that each 0 is replaced by 1, and each 1 is
 * replaced by 0. For example, inverting [0, 1, 1] results in [1, 0, 0].
 *
 * Example 1:
 *
 *
 * Input: [[1,1,0],[1,0,1],[0,0,0]]
 * Output: [[1,0,0],[0,1,0],[1,1,1]]
 * Explanation: First reverse each row: [[0,1,1],[1,0,1],[0,0,0]].
 * Then, invert the image: [[1,0,0],[0,1,0],[1,1,1]]
 *
 *
 * Example 2:
 *
 *
 * Input: [[1,1,0,0],[1,0,0,1],[0,1,1,1],[1,0,1,0]]
 * Output: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
 * Explanation: First reverse each row:
 * [[0,0,1,1],[1,0,0,1],[1,1,1,0],[0,1,0,1]].
 * Then invert the image: [[1,1,0,0],[0,1,1,0],[0,0,0,1],[1,0,1,0]]
 *
 *
 * Notes:
 *
 *
 * 1 <= A.length = A[0].length <= 20
 * 0 <= A[i][j] <= 1
 *
 *
 */

func flipAndInvertImage(A [][]int) [][]int {
	row := len(A)
	col := len(A[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col/2; j++ {
			temp := A[i][col-1-j]
			A[i][col-j-1] = A[i][j]
			A[i][j] = temp
		}
		for j := 0; j < col; j++ {
			if A[i][j] == 0 {
				A[i][j] = 1
			} else {
				A[i][j] = 0
			}
		}
	}
	return A

}

func main() {
	array := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}}
	fmt.Println("row:", len(array), "col:", len(array[0]))

	fmt.Println(flipAndInvertImage(array))
	//flipedArray := flipAndInvertImage(array)
}
