package problem890

/*
 * @lc app=leetcode id=890 lang=golang
 *
 * [890] Find and Replace Pattern
 */
//  将字符串都标准化—
func findAndReplacePattern(words []string, pattern string) []string {
	pat := normalize(pattern)

	res := []string{}
	for _, word := range words {
		wd := normalize(word)
		if wd == pat {
			res = append(res, wd)
		}
	}
	return res
}

func normalize(str string) string {
	hashmap := make(map[rune]byte, len(str))
	for _, c := range str {
		if _, ok := hashmap[c]; !ok {
			hashmap[c] = byte(len(hashmap)) + 'a'
		}
	}

	res := make([]byte, len(str))
	for i, c := range str {
		res[i] = hashmap[c]
	}
	return string(res)
}
