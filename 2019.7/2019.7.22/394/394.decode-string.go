package main

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

func decodeString(s string) string {
	n := len(s)
	// i 是数字开始下标
	i := 0
	for i < n && (s[i] < '0' || s[i] > '9') {
		i++
	}

	// 没有数字直接返回
	if i == n {
		return s
	}

	// j 是第一个[出现的下标
	j := i + 1
	for s[j] != '[' {
		j++
	}

	// k 是与 j 的 [ ,对应位置的 ]
	k := j
	count := 1
	//  count 是用来统计成对出现的[]的个数
	for count > 0 {
		k++ // 用来移动下标

		// 有新的[ 出现，增加 count 的值
		if s[k] == '[' {
			count++
		} else if s[k] == ']' {
			count--
		}
	}

	// 根据位置信息，提取数字
	num, _ := strconv.Atoi(s[i:j])

	return s[:i] + times(num, decodeString(s[j+1:k])) + decodeString(s[k+1:])
}

func times(n int, s string) string {
	res := ""
	for i := 0; i < n; i++ {
		res += s
	}
	return res
}

func main() {
	s := "2[abc]3[cd]ef"
	fmt.Println(decodeString(s))
}
