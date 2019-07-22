package problem17

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

var m = map[byte][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	res := []string{}

	backtrack(digits, 0, "", &res)
	return res
}

func backtrack(digits string, index int, cur string, res *[]string) {
	if index == len(digits) {
		*res = append(*res, cur)
		return
	}

	temp := m[digits[index]]

	for i := 0; i < len(temp); i++ {
		cur += temp[i]
		backtrack(digits, index+1, cur, res)
		//  需要复原
		cur = cur[:len(cur)-1]
	}

	// for _, v := range temp {
	// 	// cur += v
	// 	backtrack(digits, index+1, cur+v, res)
	// }

}
