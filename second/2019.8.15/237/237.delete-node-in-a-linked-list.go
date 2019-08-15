package problem237

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=237 lang=golang
 *
 * [237] Delete Node in a Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = kit.ListNode

// 这道题目比较有特点，直接对相应节点赋值处理
func deleteNode(node *ListNode) {
	next := node.Next

	// 相当于删除指定元素，使用指定节点后面的元素覆盖它
	node.Val = next.Val
	node.Next = next.Next
}
