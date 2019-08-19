package problem824

import "strings"

/*
 * @lc app=leetcode id=824 lang=golang
 *
 * [824] Goat Latin
 */
func toGoatLatin(S string) string {
	strs := strings.Split(S, " ")
	for index := range strs {
		strs[index] = handleStr(strs[index], index)
	}

	// 将数组转化为字符串的函数
	return strings.Join(strs, " ")
}

func handleStr(str string, count int) string {
	// 根据位置计算重复次数
	suffix := "ma" + strings.Repeat("a", count+1)

	if isVowe(str) {
		return str + suffix
	}
	return str[1:] + str[0:1] + suffix
}

func isVowe(str string) bool {
	switch str[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
