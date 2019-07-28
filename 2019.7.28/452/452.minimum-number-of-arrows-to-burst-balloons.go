// import "sort"

package problem452

import "sort"

/*
 * @lc app=leetcode id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */
func findMinArrowShots(points [][]int) int {

	if len(points) == 0 {
		return 0
	}

	sort.Sort(balloons(points))

	res := 1
	end := points[0][1]

	for i := 1; i < len(points); i++ {
		if points[i][0] <= end {
			continue
		}
		res++
		end = points[i][1]
	}
	return res
}

type balloons [][]int

func (b balloons) Len() int {
	return len(b)
}

func (b balloons) Less(i, j int) bool {
	return b[i][1] < b[j][1]
}

func (b balloons) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
