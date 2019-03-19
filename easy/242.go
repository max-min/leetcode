package easy

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 *
 * https://leetcode-cn.com/problems/valid-anagram/description/
 *
 * algorithms
 * Easy (50.60%)
 * Total Accepted:    23K
 * Total Submissions: 45.2K
 * Testcase Example:  '"anagram"\n"nagaram"'
 *
 * 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。
 *
 * 示例 1:
 *
 * 输入: s = "anagram", t = "nagaram"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: s = "rat", t = "car"
 * 输出: false
 *
 * 说明:
 * 你可以假设字符串只包含小写字母。
 *
 * 进阶:
 * 如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
 *
 */
func IsAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	smap := make(map[byte]int, len(s))
	tmap := make(map[byte]int, len(s))

	for i := 0; i < len(s); i++ {
		if _, exist := smap[s[i]]; exist {
			smap[s[i]]++
		} else {
			smap[s[i]] = 1
		}

		if _, exist := tmap[t[i]]; exist {
			tmap[t[i]]++
		} else {
			tmap[t[i]] = 1
		}
	}

	for k := range smap {
		if smap[k] != tmap[k] {
			return false
		}
	}
	return true
}
