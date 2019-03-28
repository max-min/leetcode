package easy

/*
 * @lc app=leetcode.cn id=459 lang=golang
 *
 * [459] 重复的子字符串
 *
 * https://leetcode-cn.com/problems/repeated-substring-pattern/description/
 *
 * algorithms
 * Easy (39.26%)
 * Total Accepted:    4.2K
 * Total Submissions: 10.7K
 * Testcase Example:  '"abab"'
 *
 * 给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。
 *
 * 示例 1:
 *
 *
 * 输入: "abab"
 *
 * 输出: True
 *
 * 解释: 可由子字符串 "ab" 重复两次构成。
 *
 *
 * 示例 2:
 *
 *
 * 输入: "aba"
 *
 * 输出: False
 *
 *
 * 示例 3:
 *
 *
 * 输入: "abcabcabcabc"
 *
 * 输出: True
 *
 * 解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)
 *
 *
 */
func RepeatedSubstringPattern(s string) bool {

	vMap := map[byte]int{}

	for i := range s {
		if _, exist := vMap[s[i]]; exist {
			vMap[s[i]]++
		} else {
			vMap[s[i]] = 1
		}
	}
	if len(vMap) == 1 {
		return true
	}

	index := 0
	first := true
	for _, v := range vMap {
		if first {
			index = v
			first = false
		}
		if index != v {
			return false
		}
	}

	return true
}
