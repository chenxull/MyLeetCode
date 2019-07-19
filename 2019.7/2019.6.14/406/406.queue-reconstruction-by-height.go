package problem406

import (
	"sort"
)

/*
 * @lc app=leetcode id=406 lang=golang
 *
 * [406] Queue Reconstruction by Height
 */
// 涉及到对二维素组的操作，需要加强联系

func reconstructQueue(people [][]int) [][]int {
	res := make([][]int, 0, len(people))

	sort.Sort(persons(people))

	insert := func(idx int, person []int) {
		res = append(res, person)
		if len(res)-1 == idx {
			return
		}
		copy(res[idx+1:], res[idx:])
		res[idx] = person
	}

	for _, p := range people {
		insert(p[1], p)
	}
	return res
}

type persons [][]int

func (p persons) Len() int { return len(p) }

// h 以降序为主
// k 以升序为辅
func (p persons) Less(i, j int) bool {
	if p[i][0] == p[j][0] {
		// 当 i 的 k 小于 j 的 k 时，返回 true，不用交换。反之返回 false 进行交换
		return p[i][1] < p[j][1]
	}
	// 当i 中的 h 大于 j 中的 h 时，返回 true，不用交换
	return p[i][0] > p[j][0]
}

func (p persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
