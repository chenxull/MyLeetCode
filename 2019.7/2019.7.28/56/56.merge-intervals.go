package problem56

import "sort"

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	// 按start 数字进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0, len(intervals))
	temp := intervals[0]

	for i := 1; i < len(intervals); i++ {
		//  cur.start <= last.end
		if intervals[i][0] <= temp[1] {
			temp[1] = max(temp[1], intervals[i][1])
		} else {
			res = append(res, temp)
			temp = intervals[i]
		}
	}
	res = append(res, temp)
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
