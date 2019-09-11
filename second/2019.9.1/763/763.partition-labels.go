package problem763

/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 */
func partitionLabels(S string) []int {
	if len(S) == 0 {
		return []int{}
	}

	hashMap := make(map[byte]int, len(S))
	for i := 0; i < len(S); i++ {
		hashMap[S[i]] = i
	}

	res := []int{}
	start, last := 0, 0
	for i := 0; i < len(S); i++ {
		// 当前部分需要遍历的最后位置
		last = max(last, hashMap[S[i]])
		if i == last {
			res = append(res, last-start+1)
			start = last + 1
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
