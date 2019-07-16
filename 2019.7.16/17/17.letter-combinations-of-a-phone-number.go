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

// 直接使用循环比那里求解，第一层循环用来遍历 获取传入的数字
//  第二层循环，控制字母的组合，目前已有的结果 res 与即将新添加的字母 一一组合
//  第三层循环，用来遍历获取 第一层中数字所代表的 字母有哪些，添加到答案中。
func letterCombinations(digits string) []string {

	res := []string{}
	if len(digits)
	// if len(digit) == 0 {
	// 	return nil
	// }
	// res := []string{}

	// for i := 0; i < len(digits); i++ {
	// 	// res 中现有的每个元素都需要在末尾加上新元素
	// 	temp := []string{}
	// 	for j := 0; j < len(res); j++ {
	// 		// 遍历数字对应的字母，并添加到 res 中
	// 		for k := 0; k < len(digit[digits[i]]); k++ {
	// 			temp = append(temp, res[j]+digit[digits[i]][k])
	// 		}
	// 	}
	// 	res = temp
	// }
	// return res
}

func helper(digits string, index int, res *[]string) {

}
