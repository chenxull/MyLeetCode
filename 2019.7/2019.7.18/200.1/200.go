package problem200

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */
 func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	if len(grid[0]) == 0 {
		return 0
	}
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			// 为避免重复计算，只有为 1 的才进入递归
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

var d = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

// 在遍历的过程中，要将数组标记，这会对数组中的元素进行修改
func dfs(grid [][]byte, r, c int) {
	// 递归终止条件

	grid[r][c] = '0'

	// 在递归前来判断将要访问的位置是否符合条件，可以减少 stack 的使用
	for i := 0; i < 4; i++ {
		newr := r + d[i][0]
		newc := c + d[i][1]
		m, n := len(grid), len(grid[0])

		if isArea(newr, newc, m, n) && grid[newr][newc] == '1' {
			dfs(grid, newr, newc)
		}
	}

}

func isArea(r, c, m, n int) bool {
	if r >= 0 && r < m && c >= 0 && c < n {
		return true
	}
	return false
}
