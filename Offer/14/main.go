package main

func integeBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1

	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[2] = max(j*(i-j), dp[j]*(i-j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {

}
