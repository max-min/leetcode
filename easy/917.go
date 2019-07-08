package easy

/*

给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。



示例 1：

输入："ab-cd"
输出："dc-ba"
示例 2：

输入："a-bC-dEf-ghIj"
输出："j-Ih-gfE-dCba"
示例 3：

输入："Test1ng-Leet=code-Q!"
输出："Qedo1ct-eeLg=ntse-T!"


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-only-letters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func ReverseOnlyLetters(S string) string {

	var letter []byte // need to do
	var chara []byte  // not change
	var index []int   // not change

	for i, s := range S {
		if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') {
			letter = append(letter, S[i])
		} else {
			chara = append(chara, S[i])
			index = append(index, i)
		}
	}

	for i, j := 0, len(letter)-1; i < j; i, j = i+1, j-1 {
		letter[i], letter[j] = letter[j], letter[i]
	}

	for i, idx := range index {
		letter = append(letter, ' ')
		copy(letter[idx+1:], letter[idx:len(letter)-1])
		letter[idx] = chara[i]
	}

	return string(letter)
}
