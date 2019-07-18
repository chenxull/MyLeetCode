package main

import "fmt"

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

func subsets(nums []int) [][]int {
	res := [][]int{}

	var cur []int
	for i := 0; i <= len(nums); i++ {
		backtracking(nums, cur, &res, 0, i)
	}
	return res
}

func backtracking(nums []int, cur []int, res *[][]int, start int, n int) {
	if len(cur) == n {
		tmp := make([]int, n)
		copy(tmp, cur)
		fmt.Println("添加元素:", tmp, "n:", n)
		*res = append(*res, tmp)
		return
	}

	//  这个 start 数值的变化 要和上一层循环中的 i 联系起来。
	//   也就是说，这一层循环的起点，是建立在上一层循环中 i 的值来决定的。这样就可以避免重复计算
	for i := start; i < len(nums); i++ {
		cur = append(cur, nums[i])
		backtracking(nums, cur, res, i+1, n)
		cur = cur[:len(cur)-1]
	}
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
