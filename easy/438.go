package easy

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 *
 * https://leetcode-cn.com/problems/find-all-anagrams-in-a-string/description/
 *
 * algorithms
 * Easy (36.01%)
 * Total Accepted:    3.7K
 * Total Submissions: 10.2K
 * Testcase Example:  '"cbaebabacd"\n"abc"'
 *
 * 给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
 *
 * 字符串只包含小写英文字母，并且字符串 s 和 p 的长度都不超过 20100。
 *
 * 说明：
 *
 *
 * 字母异位词指字母相同，但排列不同的字符串。
 * 不考虑答案输出的顺序。
 *
 *
 * 示例 1:
 *
 *
 * 输入:
 * s: "cbaebabacd" p: "abc"
 *
 * 输出:
 * [0, 6]
 *
 * 解释:
 * 起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
 * 起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * s: "abab" p: "ab"
 *
 * 输出:
 * [0, 1, 2]
 *
 * 解释:
 * 起始索引等于 0 的子串是 "ab", 它是 "ab" 的字母异位词。
 * 起始索引等于 1 的子串是 "ba", 它是 "ab" 的字母异位词。
 * 起始索引等于 2 的子串是 "ab", 它是 "ab" 的字母异位词。
 *
 *
 */

func FindAnagrams(s string, p string) []int {

	var ret []int
	if s < p {
		return ret
	} else if s == p {
		sm, pm := 0, 0
		for i := range s {
			sm += int(s[i])
			pm += int(p[i])
		}
		if sm == pm {
			ret = append(ret, 0)
		}
		return ret
	}

	pm := 0
	for j := range p {
		pm += int(p[j])
	}

	for i := 0; i < len(s); i++ {
		var sm int
		for j := i; j < len(s); j++ {
			if i+len(p) > len(s) {
				return ret
			}
			sm += int(s[j])

			if (1+j-i) == len(p) && sm == pm {
				ret = append(ret, i)
				break
			}
		}
	}
	return ret
}
