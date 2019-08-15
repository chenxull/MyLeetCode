package problem21

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = kit.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 边界情况处理
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// 直接在较小的链表上进行处理，不用申请新的链表
	// head 用来保存链表首元素，node 用来处理
	var head, node *ListNode

	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	// 遍历二个链表，进行处理
	for l1 != nil || l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}

		node = node.Next
	}

	// 连接剩余的链表数据
	if l1 != nil {
		node.Next = l1
	}

	if l2 != nil {
		node.Next = l2
	}
	return head
}
