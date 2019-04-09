package easy

/*
 * @lc app=leetcode.cn id=520 lang=golang
 *
 * [520] 检测大写字母
 *
 * https://leetcode-cn.com/problems/detect-capital/description/
 *
 * algorithms
 * Easy (51.76%)
 * Total Accepted:    6.5K
 * Total Submissions: 12.5K
 * Testcase Example:  '"USA"'
 *
 * 给定一个单词，你需要判断单词的大写使用是否正确。
 *
 * 我们定义，在以下情况时，单词的大写用法是正确的：
 *
 *
 * 全部字母都是大写，比如"USA"。
 * 单词中所有字母都不是大写，比如"leetcode"。
 * 如果单词不只含有一个字母，只有首字母大写， 比如 "Google"。
 *
 *
 * 否则，我们定义这个单词没有正确使用大写字母。
 *
 * 示例 1:
 *
 *
 * 输入: "USA"
 * 输出: True
 *
 *
 * 示例 2:
 *
 *
 * 输入: "FlaG"
 * 输出: False
 *
 *
 * 注意: 输入是由大写和小写拉丁字母组成的非空单词。
 *
 */
func DetectCapitalUse(word string) bool {
	if len(word) < 2 {
		return true
	}
	lowerword := false
	if word[0] > 96 && word[0] < 123 {
		lowerword = true
	}
	for i := 1; i < len(word); i++ {
		if lowerword && (word[i] > 64 && word[i] < 91) {
			return false
		}
		if !lowerword {
			if word[1] > 64 && word[1] < 91 {
				if word[i] < 65 || word[i] > 90 {
					return false
				}
			} else {
				if word[i] < 97 || word[i] > 122 {
					return false
				}
			}
		}

	}
	return true
}
