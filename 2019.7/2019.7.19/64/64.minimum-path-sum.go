package problem64

/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */
//   使用一维的动态规划
func minPathSum(grid [][]int) int {

	m := len(grid)
	minSum := make([][]int, m)

	n := len(grid[0])
	for i := range minSum {
		minSum[i] = make([]int, n)
	}
	//  根据当前(i,j)所在位置，来决定下次能走的位置

	minSum[0][0] = grid[0][0]

	//  先将二边的 dp 数组构建好
	// 最左边
	for i := 1; i < m; i++ {
		minSum[i][0] = grid[i][0] + minSum[i-1][0]
	}
	// 最上边
	for j := 1; j < n; j++ {
		minSum[0][j] = grid[0][j] + minSum[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			//  经过上述的处理后，就不用担心会触碰到边界条件。 (i,j) 就是所能访问到最大位置，所以不用担心越界
			minSum[i][j] = grid[i][j] + min(minSum[i-1][j], minSum[i][j-1])
		}
	}
	return minSum[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
