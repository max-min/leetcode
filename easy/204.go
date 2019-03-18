package easy

import "fmt"

/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 *
 * https://leetcode-cn.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (26.68%)
 * Total Accepted:    13.8K
 * Total Submissions: 51.6K
 * Testcase Example:  '10'
 *
 * 统计所有小于非负整数 n 的质数的数量。
 *
 * 示例:
 *
 * 输入: 10
 * 输出: 4
 * 解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 *
 *
 */
func CountPrimes(n int) int {
	var ret int
	if n < 3 {
		return 0
	}
	for n > 2 {
		n--
		t := true
		for i := 2; i < n; i++ {
			if n%i == 0 {
				t = false
				break
			}
		}
		if t {
			fmt.Println(n)
			ret++
		}

	}

	return ret

	return ret
}
