package problem491

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

type bitmap []int

func groupAnagrams(strs []string) [][]string {
	temp := make(map[int][]string, len(strs))
	//  hashtab 中相同k 的值被保存在一起
	for _, v := range strs {
		index := encode(v)
		temp[index] = append(temp[index], v)
	}

	//  如果没有加上长度参数 ，会默认生成len(temp) 个空数组，需要手动将长度设置为 0 才行
	res := make([][]string, 0, len(temp))
	//  提取分类好的值
	for _, v := range temp {
		res = append(res, v)
	}
	return res
}

var prime = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}

func encode(s string) int {
	ans := 1
	for i := 0; i < len(s); i++ {
		ans *= prime[s[i]-'a']
	}
	return ans
}
