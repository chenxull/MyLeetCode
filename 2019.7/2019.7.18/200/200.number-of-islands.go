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

	if r < 0 || c < 0 || len(grid) <= r || len(grid[0]) <= c || grid[r][c] != '1' {
		return
	}
	// 标记访问过的元素
	grid[r][c] = '0'

	// 访问四个方向的元素，这种方法更直白
	dfs(grid, r+1, c)
	dfs(grid, r-1, c)
	dfs(grid, r, c+1)
	dfs(grid, r, c-1)
}
