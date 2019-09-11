package main

import "fmt"

func main() {
	var str1 string
	var str2 string
	fmt.Scan(&str1)
	fmt.Scan(&str2)

	fmt.Println(solve(str1, str2))

}

func solve(str1, str2 string) int {
	m := len(str1)
	n := len(str2)
	// dp[][] 用来存储子问题的解
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// dp 从子问题开始求解
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			// 如果 str1 的长度为 0,唯一需要的操作就是 insert 所有 str2 的字母
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
				// 二个字符串末尾元素相同，结果和前一个状态一样
			} else if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				// 如果不相同，需要获取三种操作中，次数最少的
				dp[i][j] = 1 +
					min3(
						dp[i][j-1],   //insert
						dp[i-1][j],   // remove
						dp[i-1][j-1]) // replace
			}

		}
	}
	return dp[m][n]
}

func min3(a, b, c int) int {
	if min(a, b) < c {
		return min(a, b)
	}
	return c
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
