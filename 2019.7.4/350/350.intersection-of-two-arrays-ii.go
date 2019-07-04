package problem350

/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */
//   求二个数组的交集，需要考虑出现的次数
func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}

	m1 := getMap(nums1)
	m2 := getMap(nums2)

	// 交集肯定从较小的 map 中获取,为后续操作方便，较小的 map 存储在m1 中
	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}

	//  对 map 进行统计，计算其中每个元素出现的最小次数
	for n := range m1 {
		m1[n] = min(m1[n], m2[n])
	}

	for n, count := range m1 {
		for i := 0; i < count; i++ {
			res = append(res, n)
		}
	}
	return res
}

func getMap(n []int) map[int]int {
	res := make(map[int]int, len(n))

	//  map 的 k-v 表示为每个元素的出现次数
	for i := range n {
		res[n[i]]++
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
