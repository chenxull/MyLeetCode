// package problem63

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

	path := [100][100]int{}
	if obstacleGrid[0][0] == 0 {
		path[0][0] = 1
	}

	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 {
			path[i][0] = path[i-1][0]
		}
	}

	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 0 {
			path[0][i] = path[0][i-1]
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				path[i][j] = path[i-1][j] + path[i][j-1]
			}
		}
	}

	return path[m-1][n-1]
}
