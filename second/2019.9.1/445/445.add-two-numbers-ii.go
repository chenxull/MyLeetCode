package problem445

/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
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

// 使用二个 stack 保存链表中的数据
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1 := make([]int, 0, 128)
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}

	stack2 := make([]int, 0, 128)
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	head := &ListNode{}
	// generListNode(stack1, stack2, head)
	sum := 0
	carry := 0

	for len(stack1) != 0 || len(stack2) != 0 {
		sum = carry
		if len(stack1) != 0 {
			sum += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}

		if len(stack2) != 0 {
			sum += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}

		// 构建节点
		carry = sum / 10
		head.Val = sum % 10 //当前节点值
		front := &ListNode{Val: carry}
		front.Next = head
		head = front
	}
	if head.Val == 0 {
		return head.Next
	}
	return head

}
