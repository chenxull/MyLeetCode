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
	solution := []int{}
	backtracking(candidates, solution, target, &res)
	return res
}

func backtracking(candidates []int, solution []int, target int, res *[][]int) {
	if target == 0 {
		*res = append(*res, solution)
	}

	//  失败的返回条件
	if len(candidates) == 0 || target < candidates[0] {
		return
	}

	// 这样处理一下的用意是，让切片的容量等于长度，以后append的时候，会分配新的底层数组
	// 避免多处同时对底层数组进行修改，产生错误的答案。
	// 可以注释掉以下语句，运行单元测试，查看错误发生。
	solution = solution[:len(solution):len(solution)]

	// 递归判断
	backtracking(candidates, append(solution, candidates[0]), target-candidates[0], res)

	// 头元素不符合条件，去除继续判断
	backtracking(candidates[1:], solution, target, res)

}
