package easy

/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 *
 * https://leetcode-cn.com/problems/power-of-four/description/
 *
 * algorithms
 * Easy (44.37%)
 * Total Accepted:    6.1K
 * Total Submissions: 13.8K
 * Testcase Example:  '16'
 *
 * 给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。
 *
 * 示例 1:
 *
 * 输入: 16
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: 5
 * 输出: false
 *
 * 进阶：
 * 你能不使用循环或者递归来完成本题吗？
 *
 */
func IsPowerOfFour(num int) bool {

	/*
		if num <= 0 {
			return false
		}
		for num > 1 {
			if num%4 != 0 {
				return false
			}

			num /= 4
		}
		return true
	*/

	if num < 0 || num&(num-1) != 0 { //check(is or not) a power of 2.
		return false
	}
	if num&0x55555555 != 0 {
		return true //check 1 on odd bits

	}
	return false
}
