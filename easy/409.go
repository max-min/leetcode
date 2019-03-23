package easy

/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 *
 * https://leetcode-cn.com/problems/longest-palindrome/description/
 *
 * algorithms
 * Easy (46.50%)
 * Total Accepted:    4.2K
 * Total Submissions: 8.9K
 * Testcase Example:  '"abccccdd"'
 *
 * 给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
 *
 * 在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
 *
 * 注意:
 * 假设字符串的长度不会超过 1010。
 *
 * 示例 1:
 *
 *
 * 输入:
 * "abccccdd"
 *
 * 输出:
 * 7
 *
 * 解释:
 * 我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
 *
 *
 */
func LongestPalindrome(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	mMap := map[byte]int{}

	for i := range s {
		if _, exist := mMap[s[i]]; exist {
			mMap[s[i]]++
		} else {
			mMap[s[i]] = 1
		}
	}
	var ret int
	first := true
	for _, v := range mMap {
		if v == 1 && first {
			ret++
			first = false
		} else {
			if v%2 == 0 {
				ret += v
			} else {
				if first {
					ret++
					first = false
				}
				ret += (v - 1)
			}
		}
	}
	return ret
}
