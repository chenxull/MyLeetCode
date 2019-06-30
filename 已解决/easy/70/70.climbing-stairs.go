package problem70

/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n <= 3 {
		return n
	}
	ans := make([]int, n+1)

	ans[1] = 1
	ans[2] = 2
	for i := 3; i <= n; i++ {
		ans[i] = ans[i-2] + ans[i-1]
	}
	return ans[n]
}
