package problem39

import "sort"

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	cur := []int{}
	if len(candidates) == 0 {
		return res
	}

	helper(candidates, cur, target, &res)
	return res
}

func helper(candidates []int, cur []int, target int, res *[][]int) {
	if target == 0 {
		*res = append(*res, cur)
	}

	if len(candidates) == 0 || target < candidates[0] {
		// 返回递归的上一层
		return
	}

	cur = cur[:len(cur):len(cur)]
	//  重复添加 处于[0]为的元素
	helper(candidates, append(cur, candidates[0]), target-candidates[0], res)

	//  上述递归返回，有三种情况：  1. candidates 用完了。2.当前 candidates[0] > target 了。3.
	helper(candidates[1:], cur, target, res)
}
