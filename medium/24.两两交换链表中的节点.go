package medium

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (60.78%)
 * Total Accepted:    28.5K
 * Total Submissions: 46.9K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例:
 *
 * 给定 1->2->3->4, 你应该返回 2->1->4->3.
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SwapPairs(head *ListNode) *ListNode {

	h := head
	index := 0
	for h != nil {
		index++
		if h.Next != nil {
			index = 0
			h.Val, h.Next.Val = h.Next.Val, h.Val
			h = h.Next.Next
		} else {
			break
		}
	}

	return head
}
