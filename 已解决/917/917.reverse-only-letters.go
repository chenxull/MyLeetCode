package problem917

/*
 * @lc app=leetcode id=917 lang=golang
 *
 * [917] Reverse Only Letters
 *
 * https://leetcode.com/problems/reverse-only-letters/description/
 *
 * algorithms
 * Easy (55.91%)
 * Total Accepted:    22.2K
 * Total Submissions: 39.7K
 * Testcase Example:  '"ab-cd"'
 *
 * Given a string S, return the "reversed" string where all characters that are
 * not a letter stay in the same place, and all letters reverse their
 * positions.
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "ab-cd"
 * Output: "dc-ba"
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "a-bC-dEf-ghIj"
 * Output: "j-Ih-gfE-dCba"
 *
 *
 *
 * Example 3:
 *
 *
 * Input: "Test1ng-Leet=code-Q!"
 * Output: "Qedo1ct-eeLg=ntse-T!"
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * S.length <= 100
 * 33 <= S[i].ASCIIcode <= 122
 * S doesn't contain \ or "
 *
 *
 *
 *
 *
 */
func reverseOnlyLetters(S string) string {
	world := []byte(S)
	left := 0
	right := len(world) - 1
	for left < right {
		for left < right && !isLetter(world[left]) {
			left++
		}
		for left < right && !isLetter(world[right]) {
			right--
		}
		world[left], world[right] = world[right], world[left]
		left++
		right--
	}
	return string(world)

}

func isLetter(s byte) bool {
	return s <= 'z' && s >= 'a' ||
		s <= 'Z' && s >= 'A'
}
