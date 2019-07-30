package problem105

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	if len(inorder) == 1 {
		return root
	}
	//  根据根节点，在 inorder 数组中找到树 左右子树的 切分点
	index := findindex(root.Val, inorder)
	// 左子树的构建
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	// 右子树的构建
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}

func findindex(targt int, nums []int) int {
	for i, v := range nums {
		if targt == v {
			return i
		}
	}
	return 0
}
