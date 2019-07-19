package problem62

/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */
func uniquePaths(m int, n int) int {
	//  用来存储每个位置的路径数量
	path := [100][100]int{}

	//  想要到达第一排和第一列的路径只有一条
	for i := 0; i < m; i++ {
		path[i][0] = 1
	}
	for j := 0; j < n; j++ {
		path[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			path[i][j] = path[i-1][j] + path[i][j-1]
		}
	}
	return path[m-1][n-1]
}
