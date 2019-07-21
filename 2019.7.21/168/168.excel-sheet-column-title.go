package main

import "fmt"

/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */
func convertToTitle(n int) string {
	res := ""
	for n > 0 {
		n--
		res = string(byte(n%26)+'A') + res
		n /= 26
	}
	return res
}

func main() {
	fmt.Println(convertToTitle(701))
}
