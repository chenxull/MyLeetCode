package main

import "fmt"

// package main

// import "fmt"

/*
 * @lc app=leetcode id=926 lang=golang
 *
 * [926] Flip String to Monotone Increasing
 */
func minFlipsMonoIncr(S string) int {
	size := len(S)

	if size == 0 {
		return 0
	}
	// l[i] 用来表示将 i 位置左边的所有元素置为 0 所需要的操作次数
	// r[i] 用来表示将 i 位置右边的所有元素置为 1 所需要的操作次数
	l := make([]byte, size)
	r := make([]byte, size+1)

	l[0] = S[0] - '0'
	for i := 1; i < size; i++ {
		// 计算将 i 左边的元素全部置为 0 需要操作多少次
		l[i] = l[i-1] + S[i] - '0'
		fmt.Println("l[i]:", i, l[i])
	}

	r[size-1] = '1' - S[size-1]
	for j := size - 2; j >= 0; j-- {
		// 如果 j 位置上为 1 ，操作次数不变，如果为 0 操作次数加一
		r[j] = r[j+1] - S[j] + '1'
		fmt.Println("r[j]:", j, r[j])
	}

	ans := r[0]
	// 重头遍历整个字符串，计算出最小值
	for i := 1; i <= size; i++ {
		fmt.Println("i:", i, "l[i-1]+r[i-1]:", l[i-1]+r[i-1])
		ans = min(ans, l[i-1]+r[i])
	}
	return int(ans)
}

func min(a, b byte) byte {
	if a < b {
		return a
	}
	return b
}

func main() {
	s := "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000"
	fmt.Println(minFlipsMonoIncr(s))
}
