package problem118

// package main

/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */
func generate(numRows int) [][]int {
	res := [][]int{}
	if numRows == 0 {
		return [][]int{}
	}

	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}

	for i := 1; i < numRows; i++ {
		res = append(res, nextLevel(res[i-1]))
	}
	return res
}

func nextLevel(p []int) []int {
	//  数组扩容一格
	res := make([]int, 1, len(p)+1)
	// 将之前的转存
	res = append(res, p...)

	//  算法核心
	for i := 0; i < len(res)-1; i++ {
		res[i] += res[i+1]
	}
	return res
}
