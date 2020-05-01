package main

import "fmt"

// 将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//示例：
//
//输入：1->2->4, 1->3->4
//输出：1->1->2->3->4->4
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	newL := new(ListNode)
	nextL := newL
	for {
		if l1 == nil {
			nextL.Next = l2
			break
		}
		if l2 == nil {
			nextL.Next = l1
			break
		}
		if l1.Val < l2.Val {
			nextL.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		} else {
			nextL.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		}
		nextL = nextL.Next
	}
	return newL.Next
}

func main() {
	listNode := mergeTwoLists(
		&ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		&ListNode{1, &ListNode{3, nil}},
	)
	for listNode != nil {
		fmt.Println(listNode.Val)
		listNode = listNode.Next
	}
}
