package easy

/*
 * @lc app=leetcode.cn id=674 lang=golang
 *
 * [674] 最长连续递增序列
 *
 * https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence/description/
 *
 * algorithms
 * Easy (39.75%)
 * Total Accepted:    5.9K
 * Total Submissions: 14.9K
 * Testcase Example:  '[1,3,5,4,7]'
 *
 * 给定一个未经排序的整数数组，找到最长且连续的的递增序列。
 *
 * 示例 1:
 *
 *
 * 输入: [1,3,5,4,7]
 * 输出: 3
 * 解释: 最长连续递增序列是 [1,3,5], 长度为3。
 * 尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为5和7在原数组里被4隔开。
 *
 *
 * 示例 2:
 *
 *
 * 输入: [2,2,2,2,2]
 * 输出: 1
 * 解释: 最长连续递增序列是 [2], 长度为1。
 *
 *
 * 注意：数组长度不会超过10000。
 *
 */
func FindLengthOfLCIS(nums []int) int {

	var ret int

	for i := 0; i < len(nums)-1; i++ {
		max := 1
		for j := i; j < len(nums)-1; j++ {
			if nums[j] < nums[j+1] {
				max++
			} else {
				break
			}
		}
		if ret < max {
			ret = max
		}
	}
	return ret
}
