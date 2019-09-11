package problem452

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0

	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	fmt.Println(points)
	temp := points[0]
	res := 1
	for i := 1; i < len(points); i++ {
		if points[i][0] <= temp[1] {
			continue
		}
		res++
		temp = points[i]
	}
	return res
}
