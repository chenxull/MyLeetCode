package main

import "fmt"

/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */
func minimumTotal(triangle [][]int) int {
	//  minLen 表示此路径上的最小的值
	//  这是一个自下而上的 遍历
	if len(triangle) == 0 {
		return 0
	}

	minLen := make([]int, len(triangle))
	n := len(triangle) - 1
	//  将 triangle 最下的一层存入到路径中

	for i := 0; i < len(triangle); i++ {
		minLen[i] = triangle[n][i]
	}

	for layer := len(triangle) - 2; layer >= 0; layer-- {
		for i := 0; i <= layer; i++ {
			//  左边的 minLen 其实和右边的 minLen 代表的不是一个意思
			//  左边 minLen 是 layer 层中，i 位置的最短路径是多少。
			// 右边的 minLen 是 layer+1层中，i 位置的最短路径是多少
			fmt.Println("上一层minLen 的数据为:", minLen)
			minLen[i] = min(minLen[i], minLen[i+1]) + triangle[layer][i]
			fmt.Println("计算后 minLen 的值为:", minLen)
		}
	}
	return minLen[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	t := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{1, 5, 7},
		[]int{4, 1, 8, 1},
	}
	fmt.Println(minimumTotal(t))
}
