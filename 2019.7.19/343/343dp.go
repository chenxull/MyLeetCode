package package343

import "fmt"

//  使用动态规划的解法,从下往上
func integerBreakDP(n int) int {

	// mem[i] 表示将数字 i分割后，得到的最大值
	mem := make([]int, n+1)
	mem[1] = 1
	for i := 2; i <= n; i++ {
		// 求解 mem[i]
		for j := 1; j <= i-1; j++ {
			// mem[i] 的最大值
			fmt.Println(mem, i, j)
			mem[i] = max3(mem[i], j*(i-j), j*mem[i-j])
		}
	}
	return mem[n]
}
