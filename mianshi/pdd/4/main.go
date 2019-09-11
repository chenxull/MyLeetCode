package main

import "fmt"

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	nums := make([][]int, n)
	for i := 0; i < n; i++ {
		nums[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			//  temp := i * j
			nums[i][j] = (i + 1) * (j + 1)
		}
	}
	// fmt.Println(nums)

	mink := n*m - k + 1
	lo, hi := nums[0][0], nums[n-1][m-1]
	for lo < hi {
		mid := lo + (hi-lo)/2
		count := 0
		i, j := 0, m-1
		for i = 0; i < n; i++ {
			for j >= 0 && nums[i][j] > mid {
				j--
			}
			count += j + 1
		}
		if count < mink {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	fmt.Println(lo)
}
