package problem77

/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */
func combine(n int, k int) [][]int {
	res := [][]int{}
	cur := make([]int, k)

	backtrack(n, k, 0, 1, cur, &res)
	return res
}

func backtrack(n, k, index, start int, cur []int, res *[][]int) {
	if index == k {
		tmp := make([]int, k)
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := start; i <= n; i++ {
		// 将 元素 i 放到 index 位置上
		cur[index] = i
		backtrack(n, k, index+1, i+1, cur, res)
	}
}
