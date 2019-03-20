package easy

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 *
 * https://leetcode-cn.com/problems/move-zeroes/description/
 *
 * algorithms
 * Easy (52.81%)
 * Total Accepted:    37.7K
 * Total Submissions: 71.1K
 * Testcase Example:  '[0,1,0,3,12]'
 *
 * 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
 *
 * 示例:
 *
 * 输入: [0,1,0,3,12]
 * 输出: [1,3,12,0,0]
 *
 * 说明:
 *
 *
 * 必须在原数组上操作，不能拷贝额外的数组。
 * 尽量减少操作次数。
 *
 *
 */

func MoveZeroes(nums []int) {

	zeros := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			var j int
			for j = i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					for index := j; index < len(nums); index++ {
						nums[index-(j-i)] = nums[index]
					}
					break
				}

			}
			zeros = (j - i)
			for i := 0; i < zeros; i++ {
				nums[len(nums)-1-i] = 0
			}
		}
	}
}
