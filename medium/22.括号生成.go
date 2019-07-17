package medium

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode-cn.com/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (70.87%)
 * Total Accepted:    29.3K
 * Total Submissions: 41.2K
 * Testcase Example:  '3'
 *
 * 给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。
 *
 * 例如，给出 n = 3，生成结果为：
 *
 * [
 * ⁠ "((()))",
 * ⁠ "(()())",
 * ⁠ "(())()",
 * ⁠ "()(())",
 * ⁠ "()()()"
 * ]
 *
 *
 */
func GenerateParenthesis(n int) []string {

	if n == 0 {
		return []string{""}
	}
	var ret []string
	for n > 0 {
		var s string
		for i := 0; i < n; i++ {
			index := n
			for i := 0; i < index; i++ {
				s += "("
			}
			for i := 0; i < index; i++ {
				s += ")"
			}
		}
		ret = append(ret, s)
		n--
	}
	return ret

}
