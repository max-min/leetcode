package medium

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 *
 * https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (34.32%)
 * Total Accepted:    54.5K
 * Total Submissions: 158.4K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
 *
 * 示例：
 *
 * 给定一个链表: 1->2->3->4->5, 和 n = 2.
 *
 * 当删除了倒数第二个节点后，链表变为 1->2->3->5.
 *
 *
 * 说明：
 *
 * 给定的 n 保证是有效的。
 *
 * 进阶：
 *
 * 你能尝试使用一趟扫描实现吗？
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {

	len := 0
	t := head
	for t != nil {
		len++
		t = t.Next
	}
	if len < n {
		return head
	}
	len -= n
	if len == 0 {
		head = head.Next
		return head
	}
	t = head
	for t != nil {
		len--
		if len == 0 {
			t.Next = t.Next.Next
			return head
		}
		t = t.Next

	}
	return nil
}
