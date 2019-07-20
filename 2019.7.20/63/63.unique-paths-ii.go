package problem63

/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */
func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m := len(obstacleGrid)
	if m == 0 {
		return 0
	}

	n := len(obstacleGrid[0])
	if n == 0 {
		return 0
	}

	dp := [100][100]int{}

	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1

	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = dp[i-1][0]
		}
	}

	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 0 {
			dp[0][i] = dp[0][i-1]
		}
	}

	//  dp 核心
	for i := 1; i < m; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
