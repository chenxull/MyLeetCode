// package problem47
package main

import "fmt"

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

//   相比于 46 题，要考虑到对重复元素的处理
func permuteUnique(nums []int) [][]int {
	res := [][]int{}
	token := make([]bool, len(nums))
	cur := make([]int, len(nums))
	backtracking(nums, 0, cur, token, &res)
	return res
}

func backtracking(nums []int, index int, cur []int, token []bool, res *[][]int) {
	if index == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, cur)
		fmt.Println("get:", temp, "返回")
		*res = append(*res, temp)
		return
	}

	//  不能理解, 为何每次这样处理
	used := make(map[int]bool, len(nums)-index)
	fmt.Println("used 的长度为:", len(used))

	for i := 0; i < len(nums); i++ {
		fmt.Println("i为", i, "used 中的内容:", used)

		if !token[i] && !used[nums[i]] {
			used[nums[i]] = true
			fmt.Println("used标记", nums[i])
			//  将 i 位置标记为已使用
			token[i] = true
			cur[index] = nums[i]
			fmt.Println("进入递归函数")
			backtracking(nums, index+1, cur, token, res)
			fmt.Println("返回，恢复状态", i)
			token[i] = false
		}
	}
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}
