package problem78

/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := [][]int{}
	cur := []int{}
	// 用来控制元素数量
	for i := 0; i <= len(nums); i++ {
		backtrack(nums, cur, 0, i, &res)
	}
	return res
}

func backtrack(nums, cur []int, start, n int, res *[][]int) {
	if len(cur) == n {
		temp := make([]int, n)
		copy(temp, cur)
		*res = append(*res, temp)
		return
	}

	for i := start; i < len(nums); i++ {
		// 给 cur 添加新元素
		cur = append(cur, nums[i])
		// 进行回溯
		backtrack(nums, cur, i+1, n, res)
		// 退出上层回溯后，要将之前添加的元素从 cur 中删除
		cur = cur[:len(cur)-1]
	}
}
