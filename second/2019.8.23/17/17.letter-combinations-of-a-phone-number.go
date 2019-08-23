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

// 使用回溯法求解,多了一步数字向字符串的转化
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	res := []string{}
	cur := ""
	backtrack(digits, 0, cur, &res)
	return res
}

//  index 用来指向 digits 中元素的位置
func backtrack(digits string, index int, cur string, res *[]string) {
	if len(cur) == len(digits) {
		*res = append(*res, cur)
		return
	}
	// 获取当前数字代表的字符串
	temp := m[digits[index]] // []string{"a","b","c"}
	for i := 0; i < len(temp); i++ {
		cur += temp[i]
		backtrack(digits, index+1, cur, res)
		cur = cur[:len(cur)-1]
	}
}

// func main() {
// 	digits := "23"
// 	fmt.Println(letterCombinations(digits))
// }
