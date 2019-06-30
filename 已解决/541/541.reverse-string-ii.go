package problem541

/*
 * @lc app=leetcode id=541 lang=golang
 *
 * [541] Reverse String II
 *
 * https://leetcode.com/problems/reverse-string-ii/description/
 *
 * algorithms
 * Easy (44.97%)
 * Total Accepted:    57.2K
 * Total Submissions: 126.4K
 * Testcase Example:  '"abcdefg"\n2'
 *
 *
 * Given a string and an integer k, you need to reverse the first k characters
 * for every 2k characters counting from the start of the string. If there are
 * less than k characters left, reverse all of them. If there are less than 2k
 * but greater than or equal to k characters, then reverse the first k
 * characters and left the other as original.
 *
 *
 * Example:
 *
 * Input: s = "abcdefg", k = 2
 * Output: "bacdfeg"
 *
 *
 *
 * Restrictions:
 *
 * ⁠The string consists of lower English letters only.
 * ⁠Length of the given string and k will in the range [1, 10000]
 *
 */
func reverseStr(s string, k int) string {
	bytes := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		j := min(i+k, len(s))
		//fmt.Println("边界是:", j)
		reverse(bytes[i:j])
		//fmt.Println("反转后的字符串为:", string(bytes))
	}
	return string(bytes)
}

func reverse(bytes []byte) {
	i, j := 0, len(bytes)-1
	if i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
}
func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
