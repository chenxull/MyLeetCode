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
		hash := make(map[int]int, len(points))

		//  针对于 i，计算出，剩下的数组到期的距离，并统计个数
		for j := 0; j < len(points); j++ {
			if i != j {
				hash[distinct(points[i], points[j])]++
			}
		}

		for _, v := range hash {
			if v >= 2 {
				res += v * (v - 1)
			}

		}

	}
	return res
}

func distinct(x, y []int) int {
	a := (y[0] - x[0]) * (y[0] - x[0])
	b := (y[1] - x[1]) * (y[1] - x[1])
	return a + b
}
