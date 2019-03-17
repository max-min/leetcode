package easy

/*
 * @lc app=leetcode.cn id=168 lang=golang
 *
 * [168] Excel表列名称
 *
 * https://leetcode-cn.com/problems/excel-sheet-column-title/description/
 *
 * algorithms
 * Easy (30.29%)
 * Total Accepted:    5.8K
 * Total Submissions: 19.2K
 * Testcase Example:  '1'
 *
 * 给定一个正整数，返回它在 Excel 表中相对应的列名称。
 *
 * 例如，
 *
 * ⁠   1 -> A
 * ⁠   2 -> B
 * ⁠   3 -> C
 * ⁠   ...
 * ⁠   26 -> Z
 * ⁠   27 -> AA
 * ⁠   28 -> AB
 * ⁠   ...
 *
 *
 * 示例 1:
 *
 * 输入: 1
 * 输出: "A"
 *
 *
 * 示例 2:
 *
 * 输入: 28
 * 输出: "AB"
 *
 *
 * 示例 3:
 *
 * 输入: 701
 * 输出: "ZY"
 *
 *
 */
func ConvertToTitle(n int) string {

	var ret []byte
	for n > 26 {
		if n%26 == 0 {
			n = n/26 - 1
			ret = append(ret, 'Z')
		} else {
			ret = append(ret, byte(n%26+64))
			n /= 26
		}
	}
	ret = append(ret, byte(n+64))
	return string(ret)
}
