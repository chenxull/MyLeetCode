package problem94

/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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

func inorderTraversal(root *TreeNode) []int {
	var stack []*TreeNode
	res := []int{}

	cur := root
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, cur.Val)
		cur = cur.Right
	}
	return res

	// for cur := root; cur != nil; {

	// 	if cur.Left != nil {
	// 		stack = append(stack, cur)
	// 		cur = cur.Left
	// 	} else {
	// 		res = append(res, cur.Val)
	// 		if cur.Right != nil {
	// 			cur = cur.Right
	// 		} else {
	// 			if len(stack) == 0 {
	// 				break
	// 			}
	// 			cur = stack[len(stack)-1]
	// 			stack = stack[:len(stack)-1]
	// 		}
	// 	}
	// }
	// return res
}
