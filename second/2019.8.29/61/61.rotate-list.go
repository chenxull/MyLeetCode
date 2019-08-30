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

// 使用双指针
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	dummyhead := &ListNode{Next: head}
	slow, fast := dummyhead, dummyhead

	length := 0
	for length = 0; fast.Next != nil; length++ {
		fast = fast.Next
	}

	// 移动慢指针
	for i := length - k%length; i > 0; i-- {
		slow = slow.Next
	}

	fast.Next = dummyhead.Next
	dummyhead.Next = slow.Next
	slow.Next = nil
	return dummyhead.Next
}
