package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=709 lang=golang
 *
 * [709] To Lower Case
 *
 * https://leetcode.com/problems/to-lower-case/description/
 *
 * algorithms
 * Easy (76.11%)
 * Total Accepted:    87.8K
 * Total Submissions: 115K
 * Testcase Example:  '"Hello"'
 *
 * Implement function ToLowerCase() that has a string parameter str, and
 * returns the same string in lowercase.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "Hello"
 * Output: "hello"
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "here"
 * Output: "here"
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "LOVELY"
 * Output: "lovely"
 *
 *
 *
 *
 *
 */
func toLowerCase(str string) string {
	return strings.ToLower(str)
}

func main() {
	string := "hello"
	fmt.Println(toLowerCase(string))
}
