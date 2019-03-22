package easy

/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 *
 * https://leetcode-cn.com/problems/first-unique-character-in-a-string/description/
 *
 * algorithms
 * Easy (36.25%)
 * Total Accepted:    23.8K
 * Total Submissions: 64.9K
 * Testcase Example:  '"leetcode"'
 *
 * 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
 *
 * 案例:
 *
 *
 * s = "leetcode"
 * 返回 0.
 *
 * s = "loveleetcode",
 * 返回 2.
 *
 *
 *
 *
 * 注意事项：您可以假定该字符串只包含小写字母。
 *
 */

func FirstUniqChar(s string) int {

	for i := range s {
		exist := false
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				exist = true
				break
			}
		}
		if !exist {
			return i
		}
	}

	return -1
}
