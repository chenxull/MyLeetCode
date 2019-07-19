package problem145

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=145 lang=golang
 *
 * [145] Binary Tree Postorder Traversal
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

func postorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	res := []int{}

	cur := root
	//  先将根节点存储到 stack 中，不停的遍历存储右子树。 没有时，才返回访问左子树
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			res = append(res, cur.Val)
			cur = cur.Right
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = cur.Left
	}
	// 将结果反转即可
	return reverse(res)
}

func reverse(res []int) []int {
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return res
}
