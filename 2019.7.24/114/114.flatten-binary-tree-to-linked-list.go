package problem114

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

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

type TreeNode = kit.TreeNode

func flatten(root *TreeNode) {
	//  使用 stack 来求解
	stack := []*TreeNode{}

	p := root

	for p != nil || len(stack) != 0 {
		// 先将当前节点的右子树保存到 stack 中
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		// 来对左子树进行处理
		if p.Left != nil {
			//  剩余的子树全部在右子树中进行处理
			p.Right = p.Left
			p.Left = nil
		} else if len(stack) != 0 {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p.Right = temp
		}

		p = p.Right
	}
}
