package problem148

/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
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

// 各种排序算法都可以，使用插入排序效率有点低
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 将原链表分割成二个
	left, right := split(head)

	// 递归的分割处理，是一种从上向下的归并
	return merge(sortList(left), sortList(right))
}

// 使用快慢指针的思路来进行切分,没有使用 dummyhead 节点
func split(head *ListNode) (left, right *ListNode) {
	slow, fast := head, head

	// slowPre 用来指向slow 链表最后一个节点
	slowPre := slow
	// 一般都已 fast 指针为判断标准
	for fast != nil && fast.Next != nil {
		slowPre, slow = slow, slow.Next
		fast = fast.Next.Next
	}
	slowPre.Next = nil
	left, right = head, slow
	return
}

// 使用 dummyhead 节点便于操作.返回的时候
func merge(left, right *ListNode) *ListNode {
	dummyhead := &ListNode{}

	cur := dummyhead
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next, left = left, left.Next
		} else {
			cur.Next, right = right, right.Next
		}
		cur = cur.Next
	}

	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}

	return dummyhead.Next
}
