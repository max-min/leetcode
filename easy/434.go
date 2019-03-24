package easy

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 *
 * https://leetcode-cn.com/problems/number-of-segments-in-a-string/description/
 *
 * algorithms
 * Easy (29.02%)
 * Total Accepted:    4.3K
 * Total Submissions: 14.5K
 * Testcase Example:  '"Hello, my name is John"'
 *
 * 统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。
 *
 * 请注意，你可以假定字符串里不包括任何不可打印的字符。
 *
 * 示例:
 *
 * 输入: "Hello, my name is John"
 * 输出: 5
 *
 *
 */
func CountSegments(s string) int {

	if s == "" {
		return 0
	}
	strs := strings.Split(s, " ")

	var ret int
	for _, str := range strs {
		s := strings.Split(str, " ")
		if len(s) == 1 && s[0] != "" {
			ret++
		}
	}

	return ret
}
