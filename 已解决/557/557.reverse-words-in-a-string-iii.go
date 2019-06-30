package main

import (
	"bytes"
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 *
 * https://leetcode.com/problems/reverse-words-in-a-string-iii/description/
 *
 * algorithms
 * Easy (63.02%)
 * Total Accepted:    117K
 * Total Submissions: 184.5K
 * Testcase Example:  `"Let's take LeetCode contest"`
 *
 * Given a string, you need to reverse the order of characters in each word
 * within a sentence while still preserving whitespace and initial word order.
 *
 * Example 1:
 *
 * Input: "Let's take LeetCode contest"
 * Output: "s'teL ekat edoCteeL tsetnoc"
 *
 *
 *
 * Note:
 * In the string, each word is separated by single space and there will not be
 * any extra space in the string.
 *
 */
func reverseWords(s string) string {

	var buf bytes.Buffer
	sliceworld := strings.Split(s, " ")
	numSpace := strings.Count(s, " ")
	for n, world := range sliceworld {
		runes := []rune(world)
		for i := len(runes) - 1; i >= 0; i-- {
			buf.WriteString(string(runes[i]))
		}
		if n < numSpace {
			buf.WriteString(" ")
		}
	}
	return buf.String()
}

func main() {
	input := "Let's take LeetCode contest"
	fmt.Println(reverseWords(input))

}
