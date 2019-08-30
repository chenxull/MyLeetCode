package problem144

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
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	// 使用非递归的方式
	stack := []*TreeNode{}
	res := []int{}

	if root == nil {
		return res
	}
	// 在开始遍历时，要将root 节点放入到 stack 中
	stack = append(stack, root)
	for len(stack) != 0 {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, tmp.Val)

		// 如果当前节点的右节点不为空，先将其右节点放入到
		if tmp.Right != nil {
			stack = append(stack, tmp.Right)
		}

		// 如果存在，就放入 stack 中，这样可以确保有输出
		if tmp.Left != nil {
			stack = append(stack, tmp.Left)
		}
	}
	return res
}
