package problem763

import "fmt"

/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 */
//   使用 hashmap 来维护每个元素出现的最后位置
func partitionLabels(S string) []int {
	n := len(S)
	if n == 0 {
		return []int{}
	}

	hashmap := make(map[byte]int, n)
	res := []int{}
	for i := 0; i < n; i++ {
		hashmap[S[i]] = i
	}

	fmt.Println(hashmap)
	// start,last 分别代表，每个部分开始的位置，和结束的位置
	start, last := 0, 0
	for i := 0; i < n; i++ {
		// fmt.Println("last1:", last)
		last = max(hashmap[S[i]], last)
		// fmt.Println("last2:", last)

		if i == last {
			res = append(res, last-start+1)
			start = i + 1
		}
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
