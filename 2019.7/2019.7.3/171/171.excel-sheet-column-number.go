package problem171

/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */
func titleToNumber(s string) int {

	if len(s) == 0 {
		return 0
	}
	//  用来对字母进行转换
	hashmap := [26]int{}
	for i := 0; i < 26; i++ {
		hashmap[i] = i + 1
	}

	ans := 0
	for i := 0; i < len(s); i++ {

		ans = ans*26 + hashmap[s[i]-65]

	}

	return ans
}
