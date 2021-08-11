package main

import (
	"fmt"
)

// 题目链接：https://leetcode.com/problems/remove-nth-node-from-end-of-list/
// 题目描述：给定一个链表，要求移除倒数第n个元素，并且返回新链表的head
// 解题思路：先利用 step1 遍历n次，那么 step1 剩下的步数就是需要正数的步数了；然后利用 step2 和 step1 进行遍历即可

type listNode struct {
	Val  int
	Next *listNode
}

func main() {
	var l1 *listNode
	var n int

	// 预期结果：[1, 2, 3, 5]
	l1 = initList([]int{1, 2, 3, 4, 5})
	n = 2
	printList(removeNthFromEnd(l1, n))

	// 预期结果：[]
	l1 = initList([]int{1})
	n = 1
	printList(removeNthFromEnd(l1, n))

	// 预期结果：[1]
	l1 = initList([]int{1, 2})
	n = 1
	printList(removeNthFromEnd(l1, n))
}

func removeNthFromEnd(l1 *listNode, n int) *listNode {
	// 利用 step1 循环 n 次，往前跑 n 步（那么剩下的步数就是正数的步数）
	step1 := new(listNode)
	step1 = l1
	for i := 0; i < n; i++ {
		step1 = step1.Next
		if step1 == nil {
			return l1.Next
		}
	}

	// 这里将 step1 剩下的步数遍历完（剩下的步数就是正数的步数了）
	step2 := new(listNode)
	step2 = l1
	for step1.Next != nil {
		step1 = step1.Next
		step2 = step2.Next
	}

	// 遍历退出后，则表示找到该位置，这里进行删除即可
	if step2.Next != nil {
		step2.Next = step2.Next.Next
	}

	return l1
}

func initList(list []int) *listNode {
	if len(list) <= 0 {
		return nil
	}

	head := &listNode{Val: list[0]}
	temp := head
	for i := 1; i < len(list); i++ {
		temp.Next = &listNode{Val: list[i]}
		temp = temp.Next
	}

	return head
}

func printList(list *listNode) {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
	fmt.Printf("\n")
}
