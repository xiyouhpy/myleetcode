package main

import (
	"fmt"
)

// 题目链接：https://leetcode.com/problems/merge-two-sorted-lists/
// 题目描述：将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的
// 解题思路：分别对两个链表进行比较，比较时，值小的那个被存进链表，然后该链表向后一格再做比较，直至某一条链表到达尾部

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1, l2 *ListNode

	// 预期结果：[1, 1, 2, 3, 4, 4]
	l1, l2 = initListNode([]int{1, 2, 4}), initListNode([]int{1, 3, 4})
	printListNode(mergeTwoLists(l1, l2))

	// 预期结果：[]
	l1, l2 = initListNode([]int{}), initListNode([]int{})
	printListNode(mergeTwoLists(l1, l2))

	// 预期结果：[0]
	l1, l2 = initListNode([]int{}), initListNode([]int{0})
	printListNode(mergeTwoLists(l1, l2))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0}
	temp := head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			temp.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		} else {
			temp.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		}
		temp = temp.Next
	}
	if l1 == nil {
		temp.Next = l2
	}
	if l2 == nil {
		temp.Next = l1
	}

	return head.Next
}

func initListNode(list []int) *ListNode {
	if len(list) <= 0 {
		return nil
	}

	head := &ListNode{Val: list[0]}
	temp := head
	for i := 1; i < len(list); i++ {
		temp.Next = &ListNode{Val: list[i]}
		temp = temp.Next
	}

	return head
}

func printListNode(list *ListNode) {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
	fmt.Printf("\n")
}
