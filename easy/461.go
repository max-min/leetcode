package easy

/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 *
 * https://leetcode-cn.com/problems/hamming-distance/description/
 *
 * algorithms
 * Easy (68.47%)
 * Total Accepted:    15.5K
 * Total Submissions: 22.6K
 * Testcase Example:  '1\n4'
 *
 * 两个整数之间的汉明距离指的是这两个数字对应二进制位不同的位置的数目。
 *
 * 给出两个整数 x 和 y，计算它们之间的汉明距离。
 *
 * 注意：
 * 0 ≤ x, y < 2^31.
 *
 * 示例:
 *
 *
 * 输入: x = 1, y = 4
 *
 * 输出: 2
 *
 * 解释:
 * 1   (0 0 0 1)
 * 4   (0 1 0 0)
 * ⁠      ↑   ↑
 *
 * 上面的箭头指出了对应二进制位不同的位置。
 *
 *
 */
func HammingDistance(x int, y int) int {

	var ret int
	for {
		if x == 0 && y == 0 {
			break
		}
		var yux, yuy int
		if x > 0 {
			yux = x % 2
			x = x / 2
		}
		if y > 0 {
			yuy = y % 2
			y = y / 2
		}
		if yux-yuy != 0 {
			ret++
		}
	}

	return ret
}
