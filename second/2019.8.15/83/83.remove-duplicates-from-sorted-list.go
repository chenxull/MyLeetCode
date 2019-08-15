package problem83

import (
	"github.com/aQuaYi/LeetCode-in-Go/kit"
)

/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode = kit.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	dummyhead := &ListNode{Next: head}
	pre := dummyhead
	for pre.Next != nil {
		if pre.Next.Val == pre.Next.Next.Val {
			pre.Next = pre.Next.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return dummyhead.Next

}
