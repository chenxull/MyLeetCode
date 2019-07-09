package problem349

/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */
func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	n1 := conver2map(nums1)
	n2 := conver2map(nums2)

	//  是 n1中存储的是较小的map
	if len(n1) > len(n2) {
		n1, n2 = n2, n1
	}

	// 遍历较小的 map
	for k := range n1 {
		if n2[k] {
			res = append(res, k)
		}
	}
	return res
}

func conver2map(nums []int) map[int]bool {
	res := make(map[int]bool, len(nums))

	for i := range nums {
		//  标记存在的数字，同时避免重复数字的影响
		res[nums[i]] = true
	}
	return res
}
