package problem61

/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

//  遍历找到最后一个节点的前一个节点
func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead

	n := 0
	for n = 0; fast.Next != nil; n++ {
		fast = fast.Next
	}

	for i := n - k%n; i > 0; i-- {
		slow = slow.Next
	}

	fast.Next = dummyHead.Next
	dummyHead.Next = slow.Next
	slow.Next = nil
	return dummyHead.Next
}
