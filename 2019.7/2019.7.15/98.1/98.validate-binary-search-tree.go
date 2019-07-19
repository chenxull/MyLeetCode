package problem98

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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

//  使用 stack 的方式求解，结合中序遍历的结果，后一个节点肯定是大于前一个节点的
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}

	var stack []*TreeNode
	pre := &TreeNode{}
	for root != nil || len(stack) != 0 {
		//  中序遍历，现将左节点放入 stack 中
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pre != nil && root.Val <= pre.Val {
			return false
		}
		pre = root
		root = root.Right
	}
	return true

}
