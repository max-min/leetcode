package easy

import "strconv"

/*
 * @lc app=leetcode.cn id=415 lang=golang
 *
 * [415] 字符串相加
 *
 * https://leetcode-cn.com/problems/add-strings/description/
 *
 * algorithms
 * Easy (43.55%)
 * Total Accepted:    5.5K
 * Total Submissions: 12.5K
 * Testcase Example:  '"0"\n"0"'
 *
 * 给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。
 *
 * 注意：
 *
 *
 * num1 和num2 的长度都小于 5100.
 * num1 和num2 都只包含数字 0-9.
 * num1 和num2 都不包含任何前导零。
 * 你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式。
 *
 *
 */

func AddStrings(num1 string, num2 string) string {

	len1 := len(num1)
	len2 := len(num2)

	if len1 == 0 {
		return num2
	}
	if len2 == 0 {
		return num1
	}

	index := len1
	if len1 > len2 {
		index = len2
	}

	var ret []string
	add := false
	for i := 0; i < index; i++ {
		n := 0
		if add {
			n = 1
			add = false
		}
		n1, _ := strconv.Atoi(string(num1[len1-1-i]))
		n2, _ := strconv.Atoi(string(num2[len2-1-i]))
		n += (n1 + n2)
		if n >= 10 {
			add = true
			n %= 10
		}
		ret = append(ret, strconv.Itoa(n))
	}

	if index == len1 {
		ret = append(ret, num2[:len(num2)-index])
	} else {
		ret = append(ret, num1[:len(num1)-index])
	}

	var st string
	for i := 0; i < len(ret); i++ {
		st += ret[len(ret)-1-i]
	}
	return st
}
