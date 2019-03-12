package easy

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 *
 * https://leetcode-cn.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (28.91%)
 * Total Accepted:    19.1K
 * Total Submissions: 66.2K
 * Testcase Example:  '"Hello World"'
 *
 * 给定一个仅包含大小写字母和空格 ' ' 的字符串，返回其最后一个单词的长度。
 *
 * 如果不存在最后一个单词，请返回 0 。
 *
 * 说明：一个单词是指由字母组成，但不包含任何空格的字符串。
 *
 * 示例:
 *
 * 输入: "Hello World"
 * 输出: 5
 *
 *
 */
func LengthOfLastWord(s string) int {

	if len(s) < 1 {
		return 0
	}
	strs := strings.Split(s, " ")

	for i := len(strs) - 1; i >= 0; i-- {
		if len(strs[i]) > 0 && strs[i][0] != ' ' {
			return len(strs[i])
		}
	}
	return 0
}
