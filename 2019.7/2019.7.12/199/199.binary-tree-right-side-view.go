package problem199

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = kit.TreeNode

func rightSideView(root *TreeNode) []int {

	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	// 遍历求解，只要满足条件即可
	ls := rightSideView(root.Left)
	// 存储的只会是从根节点开始，右子树的结果
	rs := rightSideView(root.Right)

	// 左边的分支比右边长的部分，还是可以从右边看到
	if len(ls) > len(rs) {
		rs = append(rs, ls[len(rs):]...)
	}

	// 添加 root 节点
	res := make([]int, len(rs)+1)
	res[0] = root.Val
	copy(res[1:], rs)
	return res
	// res := [][]int{}
	// result := []int{}
	// var dfs func(root *TreeNode, level int)

	// dfs = func(root *TreeNode, level int) {
	// 	if root == nil {
	// 		return
	// 	}

	// 	if level >= len(res) {
	// 		res = append(res, []int{})
	// 	}
	// 	res[level] = append(res[level], root.Val)

	// 	dfs(root.Left, level+1)
	// 	dfs(root.Right, level+1)

	// }
	// dfs(root, 0)

	// for i := 0; i < len(res); i++ {
	// 	size := len(res[i])
	// 	result = append(result, res[i][size-1])
	// }
	// return result
}
