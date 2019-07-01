package problem40

import "sort"

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	solution := []int{}
	backtracking(candidates, solution, target, &res)
	return res
}

func backtracking(candidates []int, solution []int, target int, res *[][]int) {
	if target == 0 {
		*res = append(*res, solution)
	}

	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	solution = solution[:len(solution):len(solution)]

	// 判断是否符合条件
	backtracking(candidates[1:], append(solution, candidates[0]), target-candidates[0], res)

	// 不使用 candidates[0] 就需要去除数组中所有等于 candidates[0]的值
	backtracking(next(candidates), solution, target, res)

}

//  因为数组已经排序过
func next(candidates []int) []int {
	i := 0
	for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
		i++
	}
	return candidates[i+1:]
}
