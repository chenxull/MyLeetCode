package problem206

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = kit.ListNode

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	// 这里的 head 就是第一个节点
	for head != nil {
		// 保存后续节点的信息
		temp := head.Next
		// 将当前节点的位置指定 pre
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}
