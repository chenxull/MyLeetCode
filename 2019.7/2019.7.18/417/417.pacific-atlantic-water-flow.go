package problem417

/*
 * @lc app=leetcode id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 */
//  很难的一道题，建立二个二维数组，一个用来记录 坐标，有个用来标记此位置能否流入大海
//  因为是二个方向，每个二维数组要建立二个
func pacificAtlantic(matrix [][]int) [][]int {
	res := [][]int{}

	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}

	m, n := len(matrix), len(matrix[0])

	//p[i][j] ,表示[i,j]位置的水可以流到 Pacific
	// a[i][j],表示[i,j]位置的水可以流到 Atlantic
	p, a := make([][]bool, m), make([][]bool, m)
	// 创建二维 bool 型数组，创建出来的值默认为 false
	for i := 0; i < m; i++ {
		p[i] = make([]bool, n)
		a[i] = make([]bool, n)
	}

	// pQueue 是所有能让水流到 Pacific 的点的坐标的列队
	// aQueue 是所有能让水流到 Atlantic 的点的坐标的列队
	pQueue := [][]int{}
	aQueue := [][]int{}
	// 将边界的所有元素的位置坐标存储到 pQueue 和 aQueue 中
	// 先从行开始处理,同时标记 二种类型的数组
	for i := 0; i < m; i++ {
		// 最左边一层
		p[i][0] = true
		pQueue = append(pQueue, []int{i, 0})

		// 最右边一层
		a[i][n-1] = true
		aQueue = append(aQueue, []int{i, n - 1})
	}

	// 处理列
	for j := 0; j < n; j++ {
		// 最上面一层
		p[0][j] = true
		pQueue = append(pQueue, []int{0, j})
		// 最底下一层
		a[m-1][j] = true
		aQueue = append(aQueue, []int{m - 1, j})

	}

	// 核心递归逻辑
	bfs(matrix, pQueue, p)
	bfs(matrix, aQueue, a)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] && a[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

var ds = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

//  对二种类型数组进行遍历
func bfs(matrix [][]int, queue [][]int, res [][]bool) {

	for len(queue) > 0 {
		// 提取队列中第一个元素
		c := queue[0]
		queue = queue[1:]

		// 遍历每一个元素 的上下左右位置，判断是否符合条件
		for _, d := range ds {
			newi, newj := c[0]+d[0], c[1]+d[1]

			//  判断核心
			if newi >= 0 && newj >= 0 && newi < len(matrix) && newj < len(matrix[0]) &&
				!res[newi][newj] && // 判断这个节点之前是否访问
				matrix[newi][newj] >= matrix[c[0]][c[1]] { // 能否流入的关键,也就是数值上的比较
				res[newi][newj] = true
				//  将新的节点放入队列中
				queue = append(queue, []int{newi, newj})
			}
		}
	}
}
