package problem78

import "math"

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
func subsets(nums []int) [][]int {

	res := make([][]int, 0, resSize(nums))
	rec(nums, []int{}, &res)
	return res
}

func resSize(nums []int) int {
	return int(math.Pow(2, float64(len(nums))))
}

func rec(nums, temp []int, res *[][]int) {
	size := len(nums)
	// 回溯的条件
	if size == 0 {
		*res = append(*res, temp)
		return
	}

	// 对于 last 来说，要么存在于某个子集中，要么不存在
	nums, last := nums[:size-1], nums[size-1]

	rec(nums, temp, res)

	withLast := make([]int, 1, 1+len(temp))
	withLast[0] = last
	withLast = append(withLast, temp...)
	rec(nums, withLast, res)
}
