package problem337

/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var dfs func(*TreeNode) (int, int)

	// a 抢劫 root 时，最大值
	// b 不抢劫 root 时，最大值
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		la, lb := dfs(root.Left)
		ra, rb := dfs(root.Right)

		a := root.Val + lb + rb
		b := max(la, lb) + max(ra, rb)
		return a, b
	}
	a, b := dfs(root)
	return max(a, b)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
