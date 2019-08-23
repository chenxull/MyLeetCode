package problem49

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */
var prime = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}

//  分词器，每个字母对应一个素数，这样具有相同字母的字符串对应的 数字就是一样的
func groupAnagrams(strs []string) [][]string {
	// 使用 map 保存结果 key: index value:[]string
	table := make(map[int][]string, len(strs))

	for _, str := range strs {
		index := encode(str)
		table[index] = append(table[index], str)
	}

	res := make([][]string, 0, len(table))
	for _, v := range table {
		res = append(res, v)
	}

	return res
}

// 对字符串进行编码
func encode(str string) int {
	code := 1
	for i := 0; i < len(str); i++ {
		code *= prime[str[i]-'a']
	}
	return code
}
