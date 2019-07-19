package problem347

import "sort"

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */
func topKFrequent(nums []int, k int) []int {
	res := make([]int, 0, k)

	// 统计每个数字出现的次数
	rec := make(map[int]int, len(nums))
	for _, value := range nums {
		rec[value]++
	}

	// 提取 map 中的次数值 ,按照出现的次数排序
	counts := make([]int, 0, len(rec))
	for _, c := range rec {
		counts = append(counts, c)
	}

	// 排序
	sort.Ints(counts)

	minCount := counts[len(counts)-k]

	for k, v := range rec {
		if v >= minCount {
			res = append(res, k)
		}
	}
	return res

}
