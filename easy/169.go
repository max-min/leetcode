package easy

/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 求众数
 *
 * https://leetcode-cn.com/problems/majority-element/description/
 *
 * algorithms
 * Easy (58.05%)
 * Total Accepted:    28.1K
 * Total Submissions: 48.4K
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
 *
 * 你可以假设数组是非空的，并且给定的数组总是存在众数。
 *
 * 示例 1:
 *
 * 输入: [3,2,3]
 * 输出: 3
 *
 * 示例 2:
 *
 * 输入: [2,2,1,1,1,2,2]
 * 输出: 2
 *
 *
 */
func MajorityElement(nums []int) int {
	lens := len(nums)

	if lens == 1 {
		return nums[0]
	}
	tmp := make(map[int]int, lens)
	for i := range nums {
		if _, exits := tmp[nums[i]]; exits {
			tmp[nums[i]]++

			if tmp[nums[i]] > lens/2 {
				return nums[i]
			}
		} else {
			tmp[nums[i]] = 1
		}
	}
	return 0
}
