package easy

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现strStr()
 *
 * https://leetcode-cn.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (37.89%)
 * Total Accepted:    39.5K
 * Total Submissions: 104.4K
 * Testcase Example:  '"hello"\n"ll"'
 *
 * 实现 strStr() 函数。
 *
 * 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置
 * (从0开始)。如果不存在，则返回  -1。
 *
 * 示例 1:
 *
 * 输入: haystack = "hello", needle = "ll"
 * 输出: 2
 *
 *
 * 示例 2:
 *
 * 输入: haystack = "aaaaa", needle = "bba"
 * 输出: -1
 *
 *
 * 说明:
 *
 * 当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
 *
 * 对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
 *
 */
func StrStr(haystack string, needle string) int {
	if len(needle) < 1 {
		return 0
	}
	if len(haystack) < 1 {
		return -1
	}
	var rets []int
	for j := range haystack {
		if needle[0] == haystack[j] {
			rets = append(rets, j)
		}
	}
	if len(rets) < 1 {
		return -1
	}

	for _, ret := range rets {
		val := 0
		for i := 1; i < len(needle); i++ {
			if ret+i >= len(haystack) {
				val = -1
				break
			}
			if needle[i] != haystack[ret+i] {
				val = -1
				break
			}
		}
		if val != -1 {
			return ret
		}

	}
	return -1
}
