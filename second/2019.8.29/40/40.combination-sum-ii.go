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
	backtrack(candidates, target, solution, &res)
	return res
}

func backtrack(candidates []int, target int, solution []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, solution)
	}

	if len(candidates) == 0 || candidates[0] > target {
		return
	}
	solution = solution[:len(solution):len(solution)]
	backtrack(candidates[1:], target-candidates[0], append(solution, candidates[0]), res)
	backtrack(next(candidates), target, solution, res)
}

func next(candidates []int) []int {
	i := 0
	for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
		i++
	}
	return candidates[i+1:]
}
