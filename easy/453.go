package easy

/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小移动次数使数组元素相等
 *
 * https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/description/
 *
 * algorithms
 * Easy (49.65%)
 * Total Accepted:    2.8K
 * Total Submissions: 5.6K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个长度为 n 的非空整数数组，找到让数组所有元素相等的最小移动次数。每次移动可以使 n - 1 个元素增加 1。
 *
 * 示例:
 *
 *
 * 输入:
 * [1,2,3]
 *
 * 输出:
 * 3
 *
 * 解释:
 * 只需要3次移动（注意每次移动会增加两个元素的值）：
 *
 * [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
 *
 *
 */
func MinMoves(nums []int) int {

	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	sum := 0
	index := nums[0]
	repeats := true
	for _, num := range nums {
		sum += num
		if index != num {
			repeats = false
		}
	}
	if repeats {
		return 0
	}
	ret := 1
	for {
		if (ret*(numsLen-1)+sum)%((numsLen-1)*sum) == 0 {
			return ret
		} else {
			ret++
		}
	}
}
