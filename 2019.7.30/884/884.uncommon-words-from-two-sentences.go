package problem884

import "strings"

/*
 * @lc app=leetcode id=884 lang=golang
 *
 * [884] Uncommon Words from Two Sentences
 */
func uncommonFromSentences(A string, B string) []string {
	strA := strings.Split(A, " ")
	strB := strings.Split(B, " ")
	res := make([]string, 0, len(strA)+len(strB))
	hashmap := make(map[string]int, len(strA)+len(strB))

	for _, c := range strA {
		hashmap[c]++
	}
	for _, c := range strB {
		hashmap[c]++
	}

	for k, v := range hashmap {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
