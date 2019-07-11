package problem19

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
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

// 使用快慢指针，找到待删除节点的前节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}

	headPre := &ListNode{Next: head}

	slow, fast := headPre, headPre
	count := 0
	// 与 slow 拉开 n 个距离
	for fast.Next != nil && count < n {
		fast = fast.Next
		count++
	}

	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}
	delNode := slow.Next
	slow.Next = delNode.Next

	return headPre.Next
}
