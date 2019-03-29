package easy

import "strconv"

/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 *
 * https://leetcode-cn.com/problems/base-7/description/
 *
 * algorithms
 * Easy (43.27%)
 * Total Accepted:    3K
 * Total Submissions: 6.9K
 * Testcase Example:  '100'
 *
 * 给定一个整数，将其转化为7进制，并以字符串形式输出。
 *
 * 示例 1:
 *
 *
 * 输入: 100
 * 输出: "202"
 *
 *
 * 示例 2:
 *
 *
 * 输入: -7
 * 输出: "-10"
 *
 *
 * 注意: 输入范围是 [-1e7, 1e7] 。
 *
 */
func ConvertToBase7(num int) string {

	if num == 0 {
		return "0"
	}

	flag := false
	if num < 0 {
		flag = true
	}
	if flag {
		num *= -1
	}
	var ret []string
	for num > 0 {
		ret = append(ret, strconv.Itoa(num%7))
		num /= 7
	}

	var rets string
	if flag {
		rets += "-"
	}
	for i := len(ret) - 1; i >= 0; i-- {
		rets += ret[i]
	}

	return rets
}
