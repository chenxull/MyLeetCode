package problem819

import "strings"

/*
 * @lc app=leetcode id=819 lang=golang
 *
 * [819] Most Common Word
 */
func mostCommonWord(paragraph string, banned []string) string {

	isBanned := make(map[string]bool, len(banned))
	for _, b := range banned {
		isBanned[b] = true
	}

	ss := split(paragraph)

	count := make(map[string]int, len(ss))
	for _, s := range ss {
		if isBanned[s] {
			continue
		}
		count[s]++
	}

	max := 0
	res := ""

	for s, c := range count {
		if c > max {
			max = c
			res = s
		}
	}
	return res
}

//  切分字符串

func split(paragraph string) []string {
	// 去除符号
	s := replace(paragraph)
	s = strings.ToLower(s)
	return strings.Split(s, " ")
}

func replace(s string) string {
	chars := []string{"!", "?", ",", "'", ";", "."}
	for _, c := range chars {
		s = strings.Replace(s, c, " ", -1)
	}
	return s
}
