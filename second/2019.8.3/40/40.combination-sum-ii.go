package problem40

import "sort"

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */
func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	sort.Ints(candidates)
	res := [][]int{}
	cur := []int{}
	helper(candidates, cur, target, &res)
	return res
}

func helper(candidates []int, cur []int, target int, res *[][]int) {
	if target == 0 {
		*res = append(*res, cur)
	}

	if len(candidates) == 0 || candidates[0] > target {
		return
	}

	cur = cur[:len(cur):len(cur)]

	helper(candidates[1:], append(cur, candidates[0]), target-candidates[0], res)
	helper(removeDuplicate(candidates), cur, target, res)
}

//  因为是排序的，重复的元素都在一起，从 0 开始
func removeDuplicate(candidatess []int) []int {
	i := 0
	for i+1 < len(candidatess) && candidatess[i] == candidatess[i+1] {
		i++
	}
	return candidatess[i+1:]
}
