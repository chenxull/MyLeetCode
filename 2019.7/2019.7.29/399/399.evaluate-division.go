package problem399

/*
 * @lc app=leetcode id=399 lang=golang
 *
 * [399] Evaluate Division
 */
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 对输入的数据建立映射
	m := make(map[string]map[string]float64)
	for i, e := range equations {
		a, b := e[0], e[1]
		v := values[i]

		// 建立 a/b
		if _, ok := m[a]; !ok {
			m[a] = make(map[string]float64)
		}
		// todo 有待确定
		m[a][b] = 1.0 / v
		// 建立 b/a
		if _, ok := m[b]; !ok {
			m[b] = make(map[string]float64)
		}
		m[b][a] = v
	}

	// 将输入的数据构建成类似图的形式,bfs 搜索所有的结果
	res := make([]float64, len(queries))
	for i, q := range queries {
		res[i] = bfs(m, q[0], q[1])
	}
	return res
}

type entry struct {
	s string
	f float64
}

// 广度遍历整个 map，二个 map 构成了类似于图的形式
func bfs(m map[string]map[string]float64, a, b string) float64 {
	_, ok := m[a]
	if !ok {
		return -1.0
	}

	_, ok = m[b]
	if !ok {
		return -1.0
	}

	if a == b {
		return 1.0
	}

	isVisited := make(map[string]bool)
	queue := []entry{{a, 1.0}}

	for len(queue) > 0 {
		e := queue[0]
		queue = queue[1:]
		if e.s == b {
			return 1.0 / e.f
		}

		if isVisited[e.s] {
			continue
		}

		isVisited[e.s] = true

		for k, v := range m[e.s] {
			queue = append(queue, entry{k, v * e.f})
		}

	}
	return -1.0

}
