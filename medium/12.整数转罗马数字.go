package medium

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 *
 * https://leetcode-cn.com/problems/integer-to-roman/description/
 *
 * algorithms
 * Medium (59.67%)
 * Total Accepted:    28.6K
 * Total Submissions: 47.9K
 * Testcase Example:  '3'
 *
 * 罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。
 *
 * 字符          数值
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 *
 * 例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V +
 * II 。
 *
 * 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数
 * 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
 *
 *
 * I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
 * X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
 * C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
 *
 *
 * 给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。
 *
 * 示例 1:
 *
 * 输入: 3
 * 输出: "III"
 *
 * 示例 2:
 *
 * 输入: 4
 * 输出: "IV"
 *
 * 示例 3:
 *
 * 输入: 9
 * 输出: "IX"
 *
 * 示例 4:
 *
 * 输入: 58
 * 输出: "LVIII"
 * 解释: L = 50, V = 5, III = 3.
 *
 *
 * 示例 5:
 *
 * 输入: 1994
 * 输出: "MCMXCIV"
 * 解释: M = 1000, CM = 900, XC = 90, IV = 4.
 *
 */

func ge(number int) string {

	var s string
	switch number {
	case 1:
		s = "I"
	case 2:
		s = "II"
	case 3:
		s = "III"
	case 4:
		s = "IV"
	case 5:
		s = "V"
	case 6:
		s = "VI"
	case 7:
		s = "VII"
	case 8:
		s = "VIII"
	case 9:
		s = "IX"
	default:
		s = ""
	}
	return s
}
func shi(number int) string {

	var s string
	switch number {
	case 1:
		s = "X"
	case 2:
		s = "XX"
	case 3:
		s = "XXX"
	case 4:
		s = "XL"
	case 5:
		s = "L"
	case 6:
		s = "LX"
	case 7:
		s = "LXX"
	case 8:
		s = "LXXX"
	case 9:
		s = "XC"
	default:
		s = ""
	}
	return s
}

func bai(number int) string {

	var s string
	switch number {
	case 1:
		s = "C"
	case 2:
		s = "CC"
	case 3:
		s = "CCC"
	case 4:
		s = "CD"
	case 5:
		s = "D"
	case 6:
		s = "DC"
	case 7:
		s = "DCC"
	case 8:
		s = "DCCC"
	case 9:
		s = "CM"
	default:
		s = ""
	}
	return s
}

func qian(number int) string {

	var s string
	switch number {
	case 1:
		s = "M"
	case 2:
		s = "MM"
	case 3:
		s = "MMM"
	default:
		s = ""
	}
	return s
}

func IntToRoman(num int) string {

	var s string

	s = qian(num/1000) + bai(num%1000/100) + shi(num%100/10) + ge(num%10)

	return s
}
