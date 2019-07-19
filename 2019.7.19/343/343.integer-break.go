package package343


/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */
//
//  自上而下的方法，引入记忆化搜索，也就是使用一个数组缓存 已经计算过的结果
func integerBreak(n int) int {
	if n >= 2 {
		mem := make([]int, n+1)
		return breakInteger(n, mem)
	}
	return n
}

// 至少将n 分割为二个数的乘积
func breakInteger(n int, mem []int) int {
	//  当 n ==1 时停止 分割
	if n == 1 {
		return 1
	}

	//  如果已经缓存过，就直接将结果返回
	if mem[n] != 0 {
		return mem[n]
	}

	res := -1
	// 递归分割求解的过程
	for i := 1; i <= n-1; i++ {
		//  因为至少分割二个此，所以在取结果时，可以先求出 i* (n-i)
		// 然后在递归的分割下去。
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
