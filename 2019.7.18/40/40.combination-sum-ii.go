package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */
func combinationSum2(candidates []int, target int) [][]int {
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

	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	cur = cur[:len(cur):len(cur)]

	helper(candidates[1:], append(cur, candidates[0]), target-candidates[0], res)
	// 上述刚刚使用的数字无用，需要将所有后续相同的数字去除。不然可能会导致重复解
	helper(removedu(candidates), cur, target, res)
}
func removedu(candidates []int) []int {
	i := 0
	for i+1 < len(candidates) && candidates[i] == candidates[i+1] {
		i++
	}
	return candidates[i+1:]
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	fmt.Println(combinationSum2(candidates, 8))
}
