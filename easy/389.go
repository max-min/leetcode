package easy

/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 *
 * https://leetcode-cn.com/problems/find-the-difference/description/
 *
 * algorithms
 * Easy (54.59%)
 * Total Accepted:    7.1K
 * Total Submissions: 13K
 * Testcase Example:  '"abcd"\n"abcde"'
 *
 * 给定两个字符串 s 和 t，它们只包含小写字母。
 *
 * 字符串 t 由字符串 s 随机重排，然后在随机位置添加一个字母。
 *
 * 请找出在 t 中被添加的字母。
 *
 *
 *
 * 示例:
 *
 * 输入：
 * s = "abcd"
 * t = "abcde"
 *
 * 输出：
 * e
 *
 * 解释：
 * 'e' 是那个被添加的字母。
 *
 *
 */
func FindTheDifference(s string, t string) byte {

	if len(s)+1 != len(t) {
		return ' '
	}

	for i := range t {
		exist := false
		for j := range s {
			if t[i] == s[j] {
				exist = true
				break
			}
		}
		if !exist {
			return t[i]
		}
	}
	return ' '
}
