package main

import "fmt"

func main() {
	var str1 string
	var str2 string
	fmt.Scan(&str1)
	fmt.Scan(&str2)

	m := len(str1)
	n := len(str2)

	dp := make([][]int, m+1)

	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		dp[i][0] = i
	}
	for j := 0; j < n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])

			rp := 1
			if str1[i-1] == str2[j-1] {
				rp = 0
			}
			dp[i][j] = min(dp[i][j], dp[i-1][j-1]+rp)
		}
	}
	fmt.Println(dp[m][n])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
