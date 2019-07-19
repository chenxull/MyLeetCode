package problem450

/*
 * @lc app=leetcode id=450 lang=golang
 *
 * [450] Delete Node in a BST
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

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if root.Val == key {
		if root.Left == nil && root.Right == nil {
			return nil
		} else if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// 剩下的最大节点，在当前节点左子节点的 右。。。子节点上
			max := findmax(root.Left)
			// 更改当前节点的值，相当于删除了目标节点
			root.Val = max
			// 删除用来替换的节点,被删除的节点一定是一个叶子节点
			root.Left = deleteNode(root.Left, max)
		}
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root

}

func findmax(root *TreeNode) int {
	if root.Right == nil {
		return root.Val
	}
	return findmax(root.Right)
}
