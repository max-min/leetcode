package easy

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (58.57%)
 * Total Accepted:    41.2K
 * Total Submissions: 69.9K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 *
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func cycle(cur *ListNode, next *ListNode) *ListNode {
	if cur == nil {
		return nil
	}
	if next == nil {
		return cur
	}
	if next.Next == nil {
		next.Next = cur
		cur.Next = nil
	}
	next.Next = cycle(next, next.Next)
	return next
}

func ReverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	return cycle(head, head.Next)
}
