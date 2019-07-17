// package problem131

package main

import "fmt"

/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 */
func partition(s string) [][]string {
	res := [][]string{}
	current := make([]string, 0, len(s))

	backtracking(s, 0, current, &res)

	return res
}

func backtracking(s string, index int, current []string, res *[][]string) {
	//  回溯的终止条件,将 一个解添加到res 中
	if index == len(s) {
		tmp := make([]string, len(s))
		copy(tmp, current)
		*res = append(*res, tmp)
		return
	}

	//  遍历 s 中所有的元素,使用切片的方式来获取
	for j := index; j < len(s); j++ {
		if par(s[index : j+1]) {

			fmt.Println("current:", current)
			backtracking(s, j+1, append(current, s[index:j+1]), res)
		}
	}

}

//  判断字符是否为回文串
func par(s string) bool {
	if len(s) <= 1 {
		return true
	}
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
