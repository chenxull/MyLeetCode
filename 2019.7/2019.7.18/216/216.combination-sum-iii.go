package problem216

// package main

// import "fmt"

/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */
func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	cur := []int{}
	backtracking(n, k, 1, 0, cur, &res)
	return res
}

// 回溯法求解
func backtracking(n, k, start, index int, cur []int, res *[][]int) {
	if index == k && n == 0 {
		tmp := make([]int, k)
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := start; i <= 9; i++ {
		if n-i >= 0 {
			cur = append(cur, i)
			// fmt.Println(cur, index+1)
			backtracking(n-i, k, i+1, index+1, cur, res)
			cur = cur[:len(cur)-1]
		}
	}
}

// func main() {
// 	fmt.Println(combinationSum3(4, 24))
// }
