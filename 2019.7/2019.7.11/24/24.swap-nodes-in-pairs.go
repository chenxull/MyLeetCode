package problem24

/*
 * @lc app=leetcode id=24 lang=golang
 *
 * [24] Swap Nodes in Pairs
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

func swapPairs(head *ListNode) *ListNode {
	dummyhead := ListNode{Next: head}

	p := &dummyhead
	for p.Next != nil && p.Next.Next != nil {
		node1 := p.Next
		node2 := node1.Next
		next := node2.Next

		//  交换逻辑
		node2.Next = node1
		node1.Next = next
		p.Next = node2

		p = node1

	}
	return dummyhead.Next
}
