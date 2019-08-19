package main

import (
	"bytes"
	"fmt"
	"strings"
)

// package problem557

/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */
func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	// 先对单词字符串进行切分
	sliceStr := strings.Split(s, " ")
	space := len(sliceStr) - 1
	fmt.Println(space)
	var buf bytes.Buffer
	// 对每个单词进行逆转
	for n, str := range sliceStr {
		strByte := []rune(str)
		for i := len(strByte) - 1; i >= 0; i-- {
			buf.WriteString(string(strByte[i]))
		}
		if n < space {
			buf.WriteString(" ")
		}

	}
	// 拼接单词
	return buf.String()
}

func main() {
	str := "Let's take LeetCode contest"
	fmt.Println(reverseWords(str))
}
