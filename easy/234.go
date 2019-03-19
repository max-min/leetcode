package easy

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode-cn.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (35.37%)
 * Total Accepted:    17.9K
 * Total Submissions: 50.3K
 * Testcase Example:  '[1,2]'
 *
 * 请判断一个链表是否为回文链表。
 *
 * 示例 1:
 *
 * 输入: 1->2
 * 输出: false
 *
 * 示例 2:
 *
 * 输入: 1->2->2->1
 * 输出: true
 *
 *
 * 进阶：
 * 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func IsPalindromeP(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}
	var ret []int
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}

	for i := 0; i < len(ret)/2; i++ {
		if ret[i] != ret[len(ret)-1-i] {
			return false
		}
	}
	return true
}
