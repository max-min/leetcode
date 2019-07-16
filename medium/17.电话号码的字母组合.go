package medium

import (
	"fmt"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (49.83%)
 * Total Accepted:    31K
 * Total Submissions: 62.2K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */
func LetterCombinations(digits string) []string {

	sMap := make(map[int]string, 8)

	sMap[2] = "abc"
	sMap[3] = "def"
	sMap[4] = "ghi"
	sMap[5] = "jkl"
	sMap[6] = "mno"
	sMap[7] = "pqrs"
	sMap[8] = "tuv"
	sMap[9] = "wxyz"

	var rets []string
	n, _ := strconv.Atoi(digits)

	for n > 0 {
		s := n % 10
		rets = append(rets, sMap[s])
		n /= 10
	}

	var retures []string

	fmt.Println(rets)
	for j := 0; j < len(rets); j++ {
		for i := j; i < len(rets[j]); i++ {
			var tmp []byte
			tmp = append(tmp, rets[j][i])
			//fmt.Println(tmp)
			for m := i + 1; m < len(rets); m++ {
				tmp = append(tmp, rets[m][0])
			}
			fmt.Println(tmp)
			retures = append(retures, string(tmp))
		}

	}
	return rets
}
