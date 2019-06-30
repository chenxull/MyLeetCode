package problem46

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */
// 思路还需要整理清楚
func permute(nums []int) [][]int {
	n := len(nums)
	vector := make([]int, n)
	token := make([]bool, n)
	var ans [][]int

	Makepermute(0, n, nums, vector, token, &ans)
	return ans
}

func Makepermute(cur, n int, nums, vector []int, token []bool, ans *[][]int) {

	if cur == n {
		tmp := make([]int, n)
		copy(tmp, vector)
		*ans = append(*ans, tmp)
		return
	}

	for i := 0; i < n; i++ {
		if !token[i] {
			token[i] = true
			vector[cur] = nums[i]
			Makepermute(cur+1, n, nums, vector, token, ans)
			token[i] = false
		}
	}
}
