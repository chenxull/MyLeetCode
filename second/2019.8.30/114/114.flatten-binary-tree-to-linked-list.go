package problem114

/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	p := root
	stack := []*TreeNode{}
	for p != nil || len(stack) != 0 {
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			p.Right = p.Left
			p.Left = nil
		} else if len(stack) != 0 {
			tmp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p.Right = tmp
		}
		p = p.Right
	}
}
