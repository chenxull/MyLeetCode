package problem57

import "sort"

/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */
//   go 中没有方便的 insert 操作，切片对于二维数组也没用
func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	intervals = append(intervals, newInterval)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, len(intervals))
	temp := intervals[0]

	for i := 1; i < len(intervals); i++ {
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
