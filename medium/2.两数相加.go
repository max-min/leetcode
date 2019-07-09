package medium

import "fmt"

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (33.95%)
 * Total Accepted:    153.6K
 * Total Submissions: 442.9K
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 *
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 *
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 * 示例：
 *
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
 *
 *
 */

// 大数的情况下会有问题，需要考虑字符串来处理
func result(p *ListNode) int64 {

	if p.Next == nil {
		return int64(p.Val)
	}
	return result(p.Next)*10 + int64(p.Val)
}

func listResult(res int64) *ListNode {

	if res < 10 {
		return &ListNode{Val: int(res), Next: nil}
	}
	return &ListNode{Val: int(res % 10), Next: listResult(res / 10)}

}
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	fmt.Println(result(l1))
	res := result(l1) + result(l2)
	return listResult(res)
}
