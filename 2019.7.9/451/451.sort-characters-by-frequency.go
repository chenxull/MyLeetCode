package problem451

import "sort"

/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */
func frequencySort(s string) string {
	r := [256]int{}
	for i := range s {
		r[s[i]]++
	}

	ss := make([]string, 0, len(s))

	for i := range r {
		if r[i] == 0 {
			continue
		}
		// 构建[]string 数组，每个字母重复的字母需要放在一起，根据字母出现的次数来构建
		ss = append(ss, makeString(byte(i), r[i]))
	}

	//  实现了 sort 接口，对 ss 进行排序
	sort.Sort(frequency(ss))

	res := ""

	for _, str := range ss {
		res += str
	}
	return res

}

func makeString(b byte, n int) string {
	res := make([]byte, n)
	for i := range res {
		res[i] = b
	}
	return string(res)
	// res := ""
	// for i := 0; i < n; i++ {
	// 	res += string(b)
	// }
	// return res
}

type frequency []string

func (f frequency) Len() int {
	return len(f)
}

func (f frequency) Less(i, j int) bool {

	return len(f[i]) > len(f[j])
}
func (f frequency) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
