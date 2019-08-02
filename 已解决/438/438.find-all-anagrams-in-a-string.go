package problem438

/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// 比较难
func findAnagrams(s string, p string) []int {
	var res = []int{}
	var target, word [26]int

	// 给 p 中的元素做好标识，将其转化为字符数组，对应位上为 1。abc 可以表示为1 1 1 0 0 0 ...
	for i := 0; i < len(p); i++ {
		target[p[i]-'a']++
	}

	// 用来判断是否相同，标识相同的才可以判等
	check := func(i int) {
		if word == target {
			res = append(res, i)
		}
	}

	for i := 0; i < len(s); i++ {
		// 对于 s 中的字符进行加标识处理
		word[s[i]-'a']++
		// 处理第一个字符，更具 p 的长度来进行判断
		if i == len(p)-1 {
			check(0)
		} else if i >= len(p) {
			//	将不需要处理的位置 0
			word[s[i-len(p)]-'a']--
			check(i - len(p) + 1)
		}
	}
	return res
}
