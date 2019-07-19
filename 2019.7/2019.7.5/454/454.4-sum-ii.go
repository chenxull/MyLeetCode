package problem454

/*
 * @lc app=leetcode id=454 lang=golang
 *
 * [454] 4Sum II
 */
func fourSumCount(A []int, B []int, C []int, D []int) int {
	res := 0
	n := len(A)
	sum := make(map[int]int, n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum[A[i]+B[j]]++
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res += sum[-(C[i] + D[j])]
		}
	}
	return res

}
