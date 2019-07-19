package problem345

/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */
func reverseVowels(s string) string {
	b := []byte(s)
	i, j := 0, len(s)-1

	for i < j {
		for i < j && !isVowel(b[i]) {
			i++
		}
		for i < j && !isVowel(b[j]) {
			j--
		}

		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}
func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u' || b == 'A' || b == 'E' || b == 'I' || b == 'O' || b == 'U'
}
