package main

import "fmt"

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */
//

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

	res := []string{}
	if len(digits) > 0 {
		helper(digits, 0, &res, "")
	}
	return res

}

func helper(digits string, index int, res *[]string, cur string) {
	if index == len(digits) {
		fmt.Println("Get", cur, ",return")
		*res = append(*res, cur)
		return
	}

	temp := m[digits[index]]
	for _, v := range temp {
		fmt.Println("digits[", index, "]", "use", v)
		helper(digits, index+1, res, cur+v)
	}
	fmt.Println("digits[", index, "]", "commplete,return")
	return
}

func main() {
	digits := "23"
	res := letterCombinations(digits)
	fmt.Println(res)
}
