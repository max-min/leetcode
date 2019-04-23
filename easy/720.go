package easy

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=720 lang=golang
 *
 * [720] 词典中最长的单词
 *
 * https://leetcode-cn.com/problems/longest-word-in-dictionary/description/
 *
 * algorithms
 * Easy (40.72%)
 * Total Accepted:    1.9K
 * Total Submissions: 4.7K
 * Testcase Example:  '["w","wo","wor","worl","world"]'
 *
 *
 * 给出一个字符串数组words组成的一本英语词典。从中找出最长的一个单词，该单词是由words词典中其他单词逐步添加一个字母组成。若其中有多个可行的答案，则返回答案中字典序最小的单词。
 *
 * 若无答案，则返回空字符串。
 *
 * 示例 1:
 *
 *
 * 输入:
 * words = ["w","wo","wor","worl", "world"]
 * 输出: "world"
 * 解释:
 * 单词"world"可由"w", "wo", "wor", 和 "worl"添加一个字母组成。
 *
 *
 * 示例 2:
 *
 *
 * 输入:
 * words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
 * 输出: "apple"
 * 解释:
 * "apply"和"apple"都能由词典中的单词组成。但是"apple"得字典序小于"apply"。
 *
 *
 * 注意:
 *
 *
 * 所有输入的字符串都只包含小写字母。
 * words数组长度范围为[1,1000]。
 * words[i]的长度范围为[1,30]。
 *
 *
 */
func LongestWord(words []string) string {

	idx := 0
	lens := 0
	for i := range words {
		notUse := false
		for k := 1; k < len(words[i])-1; k++ {
			exist := false
			for _, t := range words {
				if strings.Compare(words[i][0:k], t) == 0 {
					exist = true
					break
				}
			}
			if !exist {
				notUse = true
				break
			}
		}
		if !notUse {
			if lens < len(words[i]) {
				idx = i
				lens = len(words[i])
			} else if lens == len(words[i]) {
				for j := 0; j < lens; j++ {
					if words[i][j] < words[idx][j] {
						idx = i
					}
				}
			}
		}

	}

	return words[idx]
}
