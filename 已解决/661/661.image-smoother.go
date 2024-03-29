package main

import "fmt"

/*
 * @lc app=leetcode id=661 lang=golang
 *
 * [661] Image Smoother
 *
 * https://leetcode.com/problems/image-smoother/description/
 *
 * algorithms
 * Easy (47.81%)
 * Total Accepted:    32.5K
 * Total Submissions: 67.3K
 * Testcase Example:  '[[1,1,1],[1,0,1],[1,1,1]]'
 *
 * Given a 2D integer matrix M representing the gray scale of an image, you
 * need to design a smoother to make the gray scale of each cell becomes the
 * average gray scale (rounding down) of all the 8 surrounding cells and
 * itself.  If a cell has less than 8 surrounding cells, then use as many as
 * you can.
 *
 * Example 1:
 *
 * Input:
 * [[1,1,1],
 * ⁠[1,0,1],
 * ⁠[1,1,1]]
 * Output:
 * [[0, 0, 0],
 * ⁠[0, 0, 0],
 * ⁠[0, 0, 0]]
 * Explanation:
 * For the point (0,0), (0,2), (2,0), (2,2): floor(3/4) = floor(0.75) = 0
 * For the point (0,1), (1,0), (1,2), (2,1): floor(5/6) = floor(0.83333333) = 0
 * For the point (1,1): floor(8/9) = floor(0.88888889) = 0
 *
 *
 *
 * Note:
 *
 * The value in the given matrix is in the range of [0, 255].
 * The length and width of the given matrix are in the range of [1, 150].
 *
 *
 */
func imageSmoother(M [][]int) [][]int {
	if M == nil || len(M) == 0 {
		return M
	}
	row := len(M)
	col := len(M[0])
	ans := make([][]int, len(M))
	for i := 0; i < len(M); i++ {
		ans[i] = make([]int, len(M[0]))
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			count := 0
			for br := i - 1; br <= i+1; br++ {
				for bc := j - 1; bc <= j+1; bc++ {
					if br >= 0 && br < row && bc >= 0 && bc < col {
						ans[i][j] += M[br][bc]
						count++
					}
				}
			}
			ans[i][j] /= count
		}
	}
	return ans
}
func main() {
	array := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	fmt.Println(imageSmoother(array))
}
