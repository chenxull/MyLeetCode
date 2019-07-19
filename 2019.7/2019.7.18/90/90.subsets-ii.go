package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	cur := []int{}
	sort.Ints(nums)
	//  控制要求的元素个数
	for i := 0; i <= len(nums); i++ {
		backtracking(i, 0, nums, cur, &res)
	}
	return res
}

func backtracking(n, start int, nums, cur []int, res *[][]int) {
	if n == len(cur) {
		tmp := make([]int, n)
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	// 循环,start 是每次新的一层中开始的位置， 需要和上一层的建立联系
	for i := start; i < len(nums); i++ {
		//  来控制重复，当重复元素第一次出现时，可以正常操作。
		// 当重复元素不是第一次出现是，就需要跳过此处操作
		if i == start || nums[i] != nums[i-1] {
			cur = append(cur, nums[i])
			backtracking(n, i+1, nums, cur, res)
			cur = cur[:len(cur)-1]
		}

	}
}

func main() {
	nums := []int{4, 4, 4, 1, 4}
	fmt.Println(subsetsWithDup(nums))
}
