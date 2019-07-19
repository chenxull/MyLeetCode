package problem203

/*
 * @lc app=leetcode id=203 lang=golang
 *
 * [203] Remove Linked List Elements
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

func removeElements(head *ListNode, val int) *ListNode {

	//  因为需要对链表中的第一个元素进行判断，所以创建一个虚头节点，方便操作。
	headPre := ListNode{Next: head}
	temp := &headPre
	for temp.Next != nil {
		if val == temp.Next.Val {
			temp.Next = temp.Next.Next
		} else {
			temp = temp.Next
		}

	}
	return headPre.Next
}
