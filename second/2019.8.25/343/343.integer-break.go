package problem343

/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */
func integerBreak(n int) int {
	//  mem 用来记录已经求出来的结果
	if n >= 2 {
		mem := make([]int, n+1)
		return breakInteger(n, mem)
	}
	return n
}

func breakInteger(n int, mem []int) int {
	if n == 1 {
		return 1
	}

	if mem[n] != 0 {
		return mem[n]
	}
	res := -1
	for i := 1; i < n; i++ {
		res = max3(res, i*(n-i), i*breakInteger(n-i, mem))
	}
	mem[n] = res

	return res

}
func max3(a, b, c int) int {
	return max(a, max(b, c))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
