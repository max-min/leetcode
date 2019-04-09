package easy

/*
 * @lc app=leetcode.cn id=541 lang=golang
 *
 * [541] 反转字符串 II
 *
 * https://leetcode-cn.com/problems/reverse-string-ii/description/
 *
 * algorithms
 * Easy (45.26%)
 * Total Accepted:    3.9K
 * Total Submissions: 8.6K
 * Testcase Example:  '"abcdefg"\n2'
 *
 * 给定一个字符串和一个整数 k，你需要对从字符串开头算起的每个 2k 个字符的前k个字符进行反转。如果剩余少于 k
 * 个字符，则将剩余的所有全部反转。如果有小于 2k 但大于或等于 k 个字符，则反转前 k 个字符，并将剩余的字符保持原样。
 *
 * 示例:
 *
 *
 * 输入: s = "abcdefg", k = 2
 * 输出: "bacdfeg"
 *
 *
 * 要求:
 *
 *
 * 该字符串只包含小写的英文字母。
 * 给定字符串的长度和 k 在[1, 10000]范围内。
 *
 *
 */

func fanzhuang(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func ReverseStr(s string, k int) string {

	lens := len(s)
	if lens < k {
		return fanzhuang(s)
	}
	if lens >= k && lens < 2*k {
		return fanzhuang(s[:k]) + s[k:]
	}

	var ret string
	var index int
	for lens > 0 {

		ret += fanzhuang(s[index : index+k])
		ret += s[index+k : index+2*k]
		lens -= 2 * k
		index += 2 * k

		if lens >= 2*k {
			continue
		} else if lens >= k && lens < 2*k {
			ret += (fanzhuang(s[index:index+k]) + s[index+k:])
		} else {
			ret += fanzhuang(s[index:])

		}
		return ret

	}
	return ret
}
