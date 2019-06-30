package problem383

/*

 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 *
 * https://leetcode.com/problems/ransom-note/description/
 *
 * algorithms
 * Easy (49.19%)
 * Total Accepted:    107.5K
 * Total Submissions: 217K
 * Testcase Example:  '"a"\n"b"'
 *
 *
 * Given an arbitrary ransom note string and another string containing letters
 * from all the magazines, write a function that will return true if the
 * ransom
 * note can be constructed from the magazines ; otherwise, it will return
 * false.
 *
 *
 * Each letter in the magazine string can only be used once in your ransom
 * note.
 *
 *
 * Note:
 * You may assume that both strings contain only lowercase letters.
 *
 *
 *
 * canConstruct("a", "b") -> false
 * canConstruct("aa", "ab") -> false
 * canConstruct("aa", "aab") -> true
 *
 *
 */
func canConstruct(ransomNote string, magazine string) bool {
	ret := make([]int, 26)
	for _, v := range magazine {
		ret[v-'a']++
	}

	for _, v := range ransomNote {
		if ret[v-'a'] >= 1 {
			ret[v-'a']--
		} else {
			return false
		}
	}
	return true
}
