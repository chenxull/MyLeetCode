package problem17

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
//

var digit = map[byte][]string{
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
	if len(digit) == 0 {
		return nil
	}
	res := []string{}

	for i := 0; i < len(digits); i++ {
		// res 中现有的每个元素都需要在末尾加上新元素
		temp := []string{}
		for j := 0; j < len(res); j++ {
			// 遍历数字对应的字母，并添加到 res 中
			for k := 0; k < len(digit[digits[i]]); k++ {
				temp = append(temp, res[j]+digit[digits[i]][k])
			}
		}
		res = temp
	}
	return res
}
