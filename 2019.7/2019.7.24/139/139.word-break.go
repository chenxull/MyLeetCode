package problem139

import "sort"

/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */
func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 || len(wordDict) == 0 {
		return false
	}

	//  用来存放字典元素
	dict := make(map[string]bool, len(wordDict))
	// 用来存放每个元素的长度
	length := make(map[int]bool, len(wordDict))

	for _, w := range wordDict {
		dict[w] = true
		length[len(w)] = true
	}

	//  将长度放入到数组中
	sizes := make([]int, 0, len(length))
	for size := range length {
		sizes = append(sizes, size)
	}

	sort.Ints(sizes)

	// dp[i] == true，等于 wordBreak(s[:i+1], wordDict) == true
	dp := make([]bool, len(s)+1)
	n := len(s)
	dp[0] = true

	for i := 0; i < len(s); i++ {
		if !dp[i] {
			continue
		}

		for _, size := range sizes {
			if i+size <= n {
				dp[i+size] = dp[i+size] || dict[s[i:i+size]]
			}
		}
	}
	return dp[n]
}
