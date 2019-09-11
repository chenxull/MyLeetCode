package problem447

/*

 * @lc app=leetcode id=447 lang=golang
 *
 * [447] Number of Boomerangs
 */
func numberOfBoomerangs(points [][]int) int {
	if len(points) < 3 {
		return 0
	}
	res := 0
	for i := 0; i < len(points); i++ {
		hashmap := make(map[int]int, len(points))
		// 统计位置 i 的节点到其他节点距离
		for j := 0; j < len(points); j++ {
			if i != j {
				hashmap[distent(points[i], points[j])]++
			}
		}

		for _, v := range hashmap {
			if v >= 2 {
				res += v * (v - 1)
			}
		}
	}
	return res
}

func distent(x, y []int) int {
	a := (y[0] - x[0]) * (y[0] - x[0])
	b := (y[1] - x[1]) * (y[1] - x[1])
	return a + b
}
