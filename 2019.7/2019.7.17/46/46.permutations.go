package main

import "fmt"

// package problem46

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
func permute(nums []int) [][]int {
	res := [][]int{}
	tmp := make([]int, len(nums))
	//  用来标记当期位的数字是否被访问
	token := make([]bool, len(nums))

	if len(nums) > 0 {
		helper(nums, 0, tmp, token, &res)
	}
	return res
}

func helper(nums []int, index int, cur []int, token []bool, res *[][]int) {
	if len(nums) == index {
		tmp := make([]int, len(nums))
		copy(tmp, cur)
		fmt.Println("get", tmp)
		*res = append(*res, tmp)
		return
	}
	//  遍历 nums 中的元素，进行组合
	for i := 0; i < len(nums); i++ {
		// 判断此元素是否被用过
		fmt.Println("判断第", i, "个元素:", nums[i])
		if !token[i] {
			token[i] = true
			fmt.Println("成功标记第", i, "个元素:", nums[i])
			cur[index] = nums[i]
			fmt.Println("DEBUG", cur)
			helper(nums, index+1, cur, token, res)
			// 返回后就取消 i 位置数字的标记.转态恢复
			token[i] = false
		}
		fmt.Println("第", i, "个元素", nums[i], "已经被标记，遍历下一个元素", i+1)
	}
}

func main() {
	nums := []int{3, 2, 5}
	fmt.Println(permute(nums))
}
