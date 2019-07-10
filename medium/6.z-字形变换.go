package medium

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 *
 * https://leetcode-cn.com/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (42.26%)
 * Total Accepted:    40.2K
 * Total Submissions: 93K
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
 *
 * 比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
 *
 * L   C   I   R
 * E T O E S I I G
 * E   D   H   N
 *
 *
 * 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
 *
 * 请你实现这个将字符串进行指定行数变换的函数：
 *
 * string convert(string s, int numRows);
 *
 * 示例 1:
 *
 * 输入: s = "LEETCODEISHIRING", numRows = 3
 * 输出: "LCIRETOESIIGEDHN"
 *
 *
 * 示例 2:
 *
 * 输入: s = "LEETCODEISHIRING", numRows = 4
 * 输出: "LDREOEIIECIHNTSG"
 * 解释:
 *
 * L     D     R
 * E   O E   I I
 * E C   I H   N
 * T     S     G
 *
 */
func Convert(s string, numRows int) string {

	// 边界条件， 当numRows 为1 或者 s长度为1的时候
	if len(s) < 3 || numRows < 2 {
		return s
	}
	ret := make([][]byte, numRows)

	loop := numRows + (numRows - 2)
	for i := 0; i < len(s); i++ {
		l := i % loop // 每次取（ numRows +中间(numRows-首尾)） 作为一次轮训
		if l < numRows {
			ret[l] = append(ret[l], s[i])
		} else {

			index := (numRows - 1) - (l - (numRows - 1)) //
			ret[index] = append(ret[index], s[i])
		}

	}

	var res []byte
	for i := range ret {
		res = append(res, ret[i]...)
	}

	return string(res[:len(s)])

}
