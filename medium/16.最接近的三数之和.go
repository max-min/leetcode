package medium

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 *
 * https://leetcode-cn.com/problems/3sum-closest/description/
 *
 * algorithms
 * Medium (41.12%)
 * Total Accepted:    31.1K
 * Total Submissions: 75.6K
 * Testcase Example:  '[-1,2,1,-4]\n1'
 *
 * 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target
 * 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
 *
 * 例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.
 *
 * 与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
 *
 *
 */

func minf(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a

}
func ThreeSumClosest(nums []int, target int) int {

	sort.Sort(ints(nums))
	lens := len(nums)
	min := nums[lens-1] + nums[lens-2] + nums[lens-3]
	s := minf(min, target)
	for i := range nums {
		second := i + 1
		third := lens - 1

		for second < third {
			t := nums[i] + nums[second] + nums[third]

			tp := minf(t, target)
			if tp < s {
				s = tp
				min = t
			} else if t > target {
				third--
			} else {
				second++
			}

		}
	}
	return min
}
