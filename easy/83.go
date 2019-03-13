package easy

/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/
 *
 * algorithms
 * Easy (44.24%)
 * Total Accepted:    18.6K
 * Total Submissions: 42K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
 *
 * 示例 1:
 *
 * 输入: 1->1->2
 * 输出: 1->2
 *
 *
 * 示例 2:
 *
 * 输入: 1->1->2->3->3
 * 输出: 1->2->3
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func DeleteDuplicates(head *ListNode) *ListNode {
	ret := head
	tmp := ret
	for head != nil {
		if head.Next == nil {
			break
		}
		if head.Val == head.Next.Val {
			head.Next = head.Next.Next
		} else {
			tmp.Next = head.Next
			tmp = tmp.Next
			head = head.Next
		}

	}
	return ret
}
