package easy

import "strings"

/*
给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。



示例 1：

输入：S = "ab#c", T = "ad#c"
输出：true
解释：S 和 T 都会变成 “ac”。
示例 2：

输入：S = "ab##", T = "c#d#"
输出：true
解释：S 和 T 都会变成 “”。
示例 3：

输入：S = "a##c", T = "#a#c"
输出：true
解释：S 和 T 都会变成 “c”。
*/

func bComapre(S string) []byte {
	var sB []byte
	lenB := 0
	for i := range S {
		if S[i] == '#' {
			if lenB == 0 {
				continue
			} else {
				lenB--
			}
		} else {
			if len(sB) > lenB {
				sB[lenB] = S[i]

			} else {
				sB = append(sB, S[i])
			}

			lenB++
		}
	}
	return sB[:lenB]
}

func BackspaceCompare(S string, T string) bool {

	s := bComapre(S)
	t := bComapre(T)

	if strings.Compare(string(s), string(t)) == 0 {
		return true
	}

	return false
}
