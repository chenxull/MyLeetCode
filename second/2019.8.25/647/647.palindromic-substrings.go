package problem647

/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 */
func countSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count += palindromic(s, i, i)
		count += palindromic(s, i, i+1)
	}
	return count
}

func palindromic(str string, i, j int) int {
	res := 0
	for i >= 0 && j < len(str) && str[i] == str[j] {
		res++
		i--
		j++

	}
	return res
}
