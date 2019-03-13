package easy

/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (46.71%)
 * Total Accepted:    17.9K
 * Total Submissions: 38.3K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给定两个二进制字符串，返回他们的和（用二进制表示）。
 *
 * 输入为非空字符串且只包含数字 1 和 0。
 *
 * 示例 1:
 *
 * 输入: a = "11", b = "1"
 * 输出: "100"
 *
 * 示例 2:
 *
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 *
 */

func generateInt(a string) int {
	var ret int
	for _, s := range a {
		ret *= 2
		if s == '1' {
			ret += 1
		}

	}
	return ret
}

func AddBinary(a string, b string) string {

	var ret []byte
	val := generateInt(a) + generateInt(b)
	for val > 0 {
		if val%2 == 0 {
			ret = append(ret, '0')
		} else {
			ret = append(ret, '1')
		}
		val /= 2
	}
	for i := 0; i < len(ret)/2; i++ {
		ret[i], ret[len(ret)-1-i] = ret[len(ret)-1-i], ret[i]
	}
	return string(ret)
}
