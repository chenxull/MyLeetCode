package problem204

/*
 * @lc app=leetcode id=204 lang=golang
 *
 * [204] Count Primes
 */

//   求解新思路
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}

	// 用来标记合数，质素的反义词
	isComposite := make([]bool, n)
	count := n / 2
	for i := 3; i*i < n; i += 2 {

		if isComposite[i] {
			continue
		}
		for j := i * i; j < n; j += 2 * i {
			if !isComposite[j] {
				count--
				isComposite[j] = true
			}
		}
	}
	return count
}
