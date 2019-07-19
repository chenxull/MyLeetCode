package problem242

/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sr := []rune(s)
	tr := []rune(t)

	res := make(map[rune]int, len(sr))

	//  遍历完成之后，如果符合题目要求，应该所有元素都为 0
	for i := range sr {
		res[sr[i]]++
		res[tr[i]]--
	}

	for _, n := range res {
		if n != 0 {
			return false
		}
	}
	return true

	// res := [26]int{}

	// for i := 0; i < len(s); i++ {
	// 	res[s[i]-97]++
	// 	res[t[i]-97]--
	// }
	// for _, n := range res {
	// 	if n != 0 {
	// 		return false
	// 	}
	// }
	// return true

}
