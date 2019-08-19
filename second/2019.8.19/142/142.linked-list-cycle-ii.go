package problem142

/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
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

// 使用快慢指针
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	dummyhead := &ListNode{Next: head}

	slow, fast := dummyhead.Next, dummyhead.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	for slow != head {
		slow, head = slow.Next, head.Next
	}
	return slow
}
