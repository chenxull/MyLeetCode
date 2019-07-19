package problem378

/*
 * @lc app=leetcode id=378 lang=golang
 *
 * [378] Kth Smallest Element in a Sorted Matrix
 */
func kthSmallest(matrix [][]int, k int) int {

	n := len(matrix)

	lo, hi := matrix[0][0], matrix[n-1][n-1]

	for lo < hi {
		mid := lo + (hi-lo)>>1
		// 用来统计此数字在数组中的位置
		count := 0
		j := n - 1

		//  这二层循环遍历结束，可以得出此 mid 在数组中的大小位置。
		for i := 0; i < n; i++ {
			for j >= 0 && matrix[i][j] > mid {
				j--
			}
			//  在这一层中，比 mid 的小的数字有几个
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
