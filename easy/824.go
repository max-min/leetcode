package easy

import "strings"

/*

给定一个由空格分割单词的句子 S。每个单词只包含大写或小写字母。

我们要将句子转换为 “Goat Latin”（一种类似于 猪拉丁文 - Pig Latin 的虚构语言）。

山羊拉丁文的规则如下：

如果单词以元音开头（a, e, i, o, u），在单词后添加"ma"。
例如，单词"apple"变为"applema"。

如果单词以辅音字母开头（即非元音字母），移除第一个字符并将它放到末尾，之后再添加"ma"。
例如，单词"goat"变为"oatgma"。

根据单词在句子中的索引，在单词最后添加与索引相同数量的字母'a'，索引从1开始。
例如，在第一个单词后添加"a"，在第二个单词后添加"aa"，以此类推。
返回将 S 转换为山羊拉丁文后的句子。

示例 1:

输入: "I speak Goat Latin"
输出: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"

*/

func isYuan(s string) bool {
	switch s[0] {
	case 'A':
		return true
	case 'a':
		return true
	case 'e':
		return true
	case 'E':
		return true
	case 'I':
		return true
	case 'i':
		return true
	case 'o':
		return true
	case 'O':
		return true
	case 'u':
		return true
	case 'U':
		return true

	}

	return false
}
func ToGoatLatin(S string) string {

	sArr := strings.Split(S, " ")
	var ret []byte
	for i := range sArr {
		if isYuan(sArr[i]) {
			ret = append(ret, sArr[i][:]...)
			ret = append(ret, 'm', 'a')
		} else {
			if len(sArr[i]) > 1 {
				ret = append(ret, sArr[i][1:]...)
			}
			ret = append(ret, sArr[i][0], 'm', 'a')
		}

		for j := 0; j <= i; j++ {
			ret = append(ret, 'a')
		}

		ret = append(ret, ' ')
	}
	return string(ret[:len(ret)-1])
}
