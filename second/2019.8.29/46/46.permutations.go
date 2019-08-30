package problem46

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
func permute(nums []int) [][]int {
	res := [][]int{}
	temp := make([]int, len(nums))
	token := make([]bool, len(nums))
	if len(nums) > 0 {
		backtrack(nums, 0, temp, token, &res)
	}
	return res
}

func backtrack(nums []int, index int, cur []int, token []bool, res *[][]int) {
	if index == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !token[i] {
			token[i] = true
			cur[index] = nums[i]
			backtrack(nums, index+1, cur, token, res)
			token[i] = false
		}
	}
}
