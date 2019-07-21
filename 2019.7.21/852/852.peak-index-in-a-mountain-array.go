package problem852

/*
 * @lc app=leetcode id=852 lang=golang
 *
 * [852] Peak Index in a Mountain Array
 */
func peakIndexInMountainArray(A []int) int {
	l := len(A)
	if l < 3 {
		return 0
	}

	for i := 1; i < l-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			return i
		}
	}
	return 0
}
