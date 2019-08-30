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

	n := len(grid)
	m := len(grid[0])
	if m == 0 {
		return 0
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' {
				bfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

func bfs(grid [][]byte, x, y int) {
	grid[x][y] = '0'
	n := len(grid)
	m := len(grid[0])

	//  将位置信息处理一下，在存入到队列中
	code := m*x + y
	queue := []int{}
	queue = append(queue, code)
	for len(queue) != 0 {
		index := queue[0]
		queue = queue[1:]
		i := index / m
		j := index % m

		if i > 0 && grid[i-1][j] == '1' {
			queue = append(queue, (i-1)*m+j)
			grid[i-1][j] = '0'
		}
		if j > 0 && grid[i][j-1] == '1' {
			queue = append(queue, i*m+j-1)
			grid[i][j-1] = '0'
		}
		if i < n-1 && grid[i+1][j] == '1' {
			queue = append(queue, (i+1)*m+j)
			grid[i+1][j] = '0'
		}
		if j < m-1 && grid[i][j+1] == '1' {
			queue = append(queue, i*m+j+1)
			grid[i][j+1] = '0'
		}
	}
}
