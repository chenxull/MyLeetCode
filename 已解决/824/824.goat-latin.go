package Problem824

import (
	"strings"
)

/*
 * @lc app=leetcode id=824 lang=golang
 *
 * [824] Goat Latin
 *
 * https://leetcode.com/problems/goat-latin/description/
 *
 * algorithms
 * Easy (56.72%)
 * Total Accepted:    28.2K
 * Total Submissions: 49.4K
 * Testcase Example:  '"I speak Goat Latin"'
 *
 * A sentence S is given, composed of words separated by spaces. Each word
 * consists of lowercase and uppercase letters only.
 *
 * We would like to convert the sentence to "Goat Latin" (a made-up language
 * similar to Pig Latin.)
 *
 * The rules of Goat Latin are as follows:
 *
 *
 * If a word begins with a vowel (a, e, i, o, or u), append "ma" to the end of
 * the word.
 * For example, the word 'apple' becomes 'applema'.
 *
 * If a word begins with a consonant (i.e. not a vowel), remove the first
 * letter and append it to the end, then add "ma".
 * For example, the word "goat" becomes "oatgma".
 *
 * Add one letter 'a' to the end of each word per its word index in the
 * sentence, starting with 1.
 * For example, the first word gets "a" added to the end, the second word gets
 * "aa" added to the end and so on.
 *
 *
 * Return the final sentence representing the conversion from S to Goat
 * Latin.
 *
 *
 *
 * Example 1:
 *
 *
 * Input: "I speak Goat Latin"
 * Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
 *
 *
 * Example 2:
 *
 *
 * Input: "The quick brown fox jumped over the lazy dog"
 * Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa
 * hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
 *
 *
 *
 *
 * Notes:
 *
 *
 * S contains only uppercase, lowercase and spaces. Exactly one space between
 * each word.
 * 1 <= S.length <= 150.
 *
 *
 */
func toGoatLatin(S string) string {
	worlds := strings.Split(S, " ")
	for index := range worlds {
		worlds[index] = handleWorld(worlds[index], index)
	}
	return strings.Join(worlds, " ")
}
func handleWorld(world string, count int) string {
	profix := "ma" + strings.Repeat("a", count+1)
	if isBeginsWithVowel(world) {
		return world + profix
	}

	return world[1:] + world[0:1] + profix
}

func isBeginsWithVowel(world string) bool {

	switch world[0] {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
