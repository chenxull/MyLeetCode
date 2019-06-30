package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

/*
 * @lc app=leetcode id=937 lang=golang
 *
 * [937] Reorder Log Files
 *
 * https://leetcode.com/problems/reorder-log-files/description/
 *
 * algorithms
 * Easy (58.42%)
 * Total Accepted:    13.8K
 * Total Submissions: 23.1K
 * Testcase Example:  '["a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"]'
 *
 * You have an array of logs.  Each log is a space delimited string of words.
 *
 * For each log, the first word in each log is an alphanumeric identifier.
 * Then, either:
 *
 *
 * Each word after the identifier will consist only of lowercase letters,
 * or;
 * Each word after the identifier will consist only of digits.
 *
 *
 * We will call these two varieties of logs letter-logs and digit-logs.  It is
 * guaranteed that each log has at least one word after its identifier.
 *
 * Reorder the logs so that all of the letter-logs come before any digit-log.
 * The letter-logs are ordered lexicographically ignoring identifier, with the
 * identifier used in case of ties.  The digit-logs should be put in their
 * original order.
 *
 * Return the final order of the logs.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: ["a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"]
 * Output: ["g1 act car","a8 act zoo","ab1 off key dog","a1 9 2 3 1","zo4 4
 * 7"]
 *
 *
 *
 *
 * Note:
 *
 *
 * 0 <= logs.length <= 100
 * 3 <= logs[i].length <= 100
 * logs[i] is guaranteed to have an identifier, and a word after the
 * identifier.
 *
 *
 */

//不会写
func reorderLogFiles(logs []string) []string {
	count := 1
	sort.SliceStable(logs, func(i, j int) bool {
		fmt.Println("i:", i, "j:", j)
		s1 := strings.SplitN(logs[i], " ", 2)
		s2 := strings.SplitN(logs[j], " ", 2)
		f1, f2 := "0"+s1[1], "0"+s2[1]

		fmt.Println("count:", count, "f1:", string(f1[1]), "f2:", string(f2[1]))
		count++
		if unicode.IsNumber(rune(f1[1])) {
			f1 = "1"
		}
		if unicode.IsNumber(rune(f2[1])) {
			f2 = "1"
		}
		return f1 < f2
	})
	return logs

}

func main() {
	logs := []string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}
	fmt.Println(reorderLogFiles(logs))
}
