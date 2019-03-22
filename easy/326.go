package easy

/*
 * @lc app=leetcode.cn id=326 lang=golang
 *
 * [326] 3的幂
 *
 * https://leetcode-cn.com/problems/power-of-three/description/
 *
 * algorithms
 * Easy (42.78%)
 * Total Accepted:    14.2K
 * Total Submissions: 33.1K
 * Testcase Example:  '27'
 *
 * 给定一个整数，写一个函数来判断它是否是 3 的幂次方。
 *
 * 示例 1:
 *
 * 输入: 27
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: 0
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: 9
 * 输出: true
 *
 * 示例 4:
 *
 * 输入: 45
 * 输出: false
 *
 * 进阶：
 * 你能不使用循环或者递归来完成本题吗？
 *
 */
func IsPowerOfThree(n int) bool {

	if n <= 0 {
		return false
	}
	for n > 1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}
	return true
}