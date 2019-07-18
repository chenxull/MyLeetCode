package main

import "fmt"

/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */
func combine(n int, k int) [][]int {
	if n <= 0 || k <= 0 || k > n {
		return [][]int{}
	}

	res := [][]int{}
	cur := make([]int, k)
	helper(n, k, 0, 1, cur, &res)
	return res
}

// cur(n,k)问题,start 是每次遍历开始的位置，cur 存储中间结果
func helper(n, k, index int, start int, cur []int, res *[][]int) {

	if index == k {
		tmp := make([]int, k)
		copy(tmp, cur)
		fmt.Println("GET", tmp)
		*res = append(*res, tmp)
		return
	}

	//  使用剪枝优化 还有 k-index 个空位，所以在[i...n]中至少需要 k -index 个空位
	// i 最多为 n-(k-index)+1 
	for i := start; i <= n-(k-index)+1; i++ {

		cur[index] = i
		fmt.Println("cur 当前值为:", cur)
		//  i+1 之前的元素以及添加到 cur 中了，可以画出回溯树来理解
		helper(n, k, index+1, i+1, cur, res)
		// 将之前放入的元素拿出来
		fmt.Println("cur 的长度为：", len(cur))
		fmt.Println("返回之后，cur 的值为:", cur, " cur的长度为", len(cur), "下一个需要放入的值为:", i+1)
	}
}

func main() {
	n, k := 10, 2
	fmt.Println(combine(n, k))
}
