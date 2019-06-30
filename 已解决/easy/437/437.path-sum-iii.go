package problem437

import "github.com/aQuaYi/LeetCode-in-Go/kit"

/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
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

// 使用 dfs 算法，加上多层递归求解
// 对于每一个节点使用 dfs 算法，来计算其子节点的和是否为目标数字
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	cnt := 0

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		sum -= node.Val
		if sum == 0 {
			cnt++
		}

		dfs(node.Left, sum)
		dfs(node.Right, sum)

	}

	dfs(root, sum)
	return cnt + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
