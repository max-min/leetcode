package easy

import "fmt"

/*
 * @lc app=leetcode.cn id=476 lang=golang
 *
 * [476] 数字的补数
 *
 * https://leetcode-cn.com/problems/number-complement/description/
 *
 * algorithms
 * Easy (67.15%)
 * Total Accepted:    7.8K
 * Total Submissions: 11.6K
 * Testcase Example:  '5'
 *
 * 给定一个正整数，输出它的补数。补数是对该数的二进制表示取反。
 *
 * 注意:
 *
 *
 * 给定的整数保证在32位带符号整数的范围内。
 * 你可以假定二进制数不包含前导零位。
 *
 *
 * 示例 1:
 *
 *
 * 输入: 5
 * 输出: 2
 * 解释: 5的二进制表示为101（没有前导零位），其补数为010。所以你需要输出2。
 *
 *
 * 示例 2:
 *
 *
 * 输入: 1
 * 输出: 0
 * 解释: 1的二进制表示为1（没有前导零位），其补数为0。所以你需要输出0。
 *
 *
 */
func FindComplement(num int) int {

	if num == 0 {
		return 0
	}
	var ret []byte
	for num > 0 {
		s := num % 2
		num /= 2
		if num == 0 && s == 0 {
			break
		}

		if s == 0 {
			ret = append(ret, '1')
		} else {
			ret = append(ret, '0')
		}
	}

	fmt.Println(ret)
	lens := len(ret)
	var retVal int
	zeros := true
	for i := lens - 1; i >= 0; i-- {
		if zeros && ret[i] == '0' {
			continue
		}
		zeros = false
		if ret[i] == '1' {
			retVal = retVal*2 + 1
		} else {
			retVal = retVal * 2
		}
	}

	return retVal
}
