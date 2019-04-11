package easy

/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 *
 * https://leetcode-cn.com/problems/sum-of-square-numbers/description/
 *
 * algorithms
 * Easy (29.87%)
 * Total Accepted:    4.7K
 * Total Submissions: 15.9K
 * Testcase Example:  '5'
 *
 * 给定一个非负整数 c ，你要判断是否存在两个整数 a 和 b，使得 a^2 + b^2 = c。
 *
 * 示例1:
 *
 *
 * 输入: 5
 * 输出: True
 * 解释: 1 * 1 + 2 * 2 = 5
 *
 *
 *
 *
 * 示例2:
 *
 *
 * 输入: 3
 * 输出: False
 *
 *
 */
func JudgeSquareSum(c int) bool {

	var lens = c
	if c < 1000 {
		lens = 10
	} else if c < 100000 {
		lens = 100
	} else if c < 10000000 {
		lens = 1000
	} else {
		lens = lens / 10000
	}
	for i := 0; i < lens; i++ {
		for j := i; j < lens; j++ {
			if i*i+j*j == c {
				return true
			}
		}
	}
	return false
}
