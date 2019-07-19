package problem437

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

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

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	count := 0

	var dfs func(*TreeNode, int)

	dfs = func(root *TreeNode, temp int) {
		if root == nil {
			return
		}
		temp -= root.Val
		if temp == 0 {
			count++
		}

		dfs(root.Left, temp)
		dfs(root.Right, temp)
	}

	dfs(root, sum)
	// 对于树中的每个节点，都当成为根节点来对待，都遍历一遍
	return count + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}
