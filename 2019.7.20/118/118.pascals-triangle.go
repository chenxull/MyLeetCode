// package problem118
package main

import "fmt"

/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */
//   完全按照计算的思路来
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}

	res := make([][]int, numRows)
	res[0] = []int{1}

	for i := 1; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0], temp[i] = 1, 1
		for j := 0; j < len(res[i-1])-1; j++ {
			temp[j+1] = res[i-1][j] + res[i-1][j+1]
		}

		res[i] = temp
	}
	return res
}

func main() {
	fmt.Println(generate(5))
}
