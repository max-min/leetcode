package easy

/*
 * @lc app=leetcode.cn id=693 lang=golang
 *
 * [693] 交替位二进制数
 *
 * https://leetcode-cn.com/problems/binary-number-with-alternating-bits/description/
 *
 * algorithms
 * Easy (59.34%)
 * Total Accepted:    5.3K
 * Total Submissions: 9K
 * Testcase Example:  '5'
 *
 * 给定一个正整数，检查他是否为交替位二进制数：换句话说，就是他的二进制数相邻的两个位数永不相等。
 *
 * 示例 1:
 *
 *
 * 输入: 5
 * 输出: True
 * 解释:
 * 5的二进制数是: 101
 *
 *
 * 示例 2:
 *
 *
 * 输入: 7
 * 输出: False
 * 解释:
 * 7的二进制数是: 111
 *
 *
 * 示例 3:
 *
 *
 * 输入: 11
 * 输出: False
 * 解释:
 * 11的二进制数是: 1011
 *
 *
 * 示例 4:
 *
 *
 * 输入: 10
 * 输出: True
 * 解释:
 * 10的二进制数是: 1010
 *
 *
 */
func HasAlternatingBits(n int) bool {

	oneFlag := false
	zeroFlag := false
	for n > 0 {
		if oneFlag && (n%2 == 1) {
			return false

		}
		if zeroFlag && (n%2 == 0) {
			return false
		}
		if n%2 == 1 {
			zeroFlag = false
			oneFlag = true
		} else {
			zeroFlag = true
			oneFlag = false
		}
		n /= 2
	}
	return true
}
