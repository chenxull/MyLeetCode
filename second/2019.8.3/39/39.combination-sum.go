package problem39

import "sort"

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	if len(candidates) == 0 {
		return nil
	}
	cur := []int{}
	res := [][]int{}
	helper(candidates, cur, target, &res)
	return res
}

// 递归的方式求解
func helper(candidates []int, cur []int, target int, res *[][]int) {
	if target == 0 {
		*res = append(*res, cur)
	}

	if len(candidates) == 0 || candidates[0] > target {
		return
	}

	cur = cur[:len(cur):len(cur)]

	helper(candidates, append(cur, candidates[0]), target-candidates[0], res)

	// 从上一层递归中出来
	helper(candidates[1:], cur, target, res)
}
