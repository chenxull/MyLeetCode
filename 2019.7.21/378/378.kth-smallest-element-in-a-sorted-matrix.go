package problem378

/*
 * @lc app=leetcode id=378 lang=golang
 *
 * [378] Kth Smallest Element in a Sorted Matrix
 */
func kthSmallest(matrix [][]int, k int) int {

	m := len(matrix)
	lo, hi := matrix[0][0], matrix[m-1][m-1]

	//二分搜索中，常用的限制条件
	for lo < hi {
		mid := lo + (hi-lo)/2
		count := 0
		i, j := 0, m-1
		for i = 0; i < m; i++ {
			// 统计这一层中有多少个数字比 mid 大
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			count += j + 1
		}

		if count < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
