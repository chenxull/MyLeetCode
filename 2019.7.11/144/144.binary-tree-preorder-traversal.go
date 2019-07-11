package problem144

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=144 lang=golang
 *
 * [144] Binary Tree Preorder Traversal
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
//  */

type TreeNode = kit.TreeNode

func preorderTraversal(root *TreeNode) []int {
	// 用来保存右节点的值
	rightStack := []*TreeNode{}
	res := []int{}
	if root == nil {
		return res
	}

	cur := root
	rightStack = append(rightStack, root)
	for len(rightStack) != 0 {
		cur = rightStack[len(rightStack)-1]
		rightStack = rightStack[:len(rightStack)-1]
		res = append(res, cur.Val)

		if cur.Right != nil {
			rightStack = append(rightStack, cur.Right)
		}
		if cur.Left != nil {
			rightStack = append(rightStack, cur.Left)
		}
	}
	return res

	// for cur := root; cur != nil; {
	// 	res = append(res, cur.Val)

	// 	if cur.Left != nil {
	// 		if cur.Right != nil {
	// 			// 将当前节点的右子节点放入 stack 中，后续的时候还会使用到
	// 			rightStack = append(rightStack, cur.Right)
	// 		}
	// 		//  处理好当前节点之后，继续向左子节点移动
	// 		cur = cur.Left
	// 	} else {
	// 		// 当前节点左子树为空，右子树存在，指针移动到右子树中
	// 		if cur.Right != nil {
	// 			cur = cur.Right
	// 		} else {
	// 			// 当前节点左右子树都不存在，需要判断是否遍历结束,判断标准为 rightstack 是否为空。
	// 			if len(rightStack) == 0 {
	// 				break
	// 			}
	// 			// 遍历还没有结束，且当前节点左右子树都为空，需要从 stack 中取出 兄弟右节点来继续处理
	// 			cur = rightStack[len(rightStack)-1]
	// 			rightStack = rightStack[:len(rightStack)-1]
	// 		}
	// 	}
	// 	//
	// }
	// return res
}
