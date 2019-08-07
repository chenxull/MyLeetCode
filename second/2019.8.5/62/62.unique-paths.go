package problem62

/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */
//   使用动态规划来求解，先将左侧和上侧的路径数量保存好
// 然后使用二维数组 dp 来保存剩余每个位置可以通过的路径数量
func uniquePaths(m int, n int) int {
	path := [100][100]int{}
	for i := 0; i < m; i++ {
		path[i][0] = 1
	}

	for i := 0; i < n; i++ {
		path[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}
	return path[m-1][n-1]
}
