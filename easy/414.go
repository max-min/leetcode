package easy

import "fmt"

/*
 * @lc app=leetcode.cn id=414 lang=golang
 *
 * [414] 第三大的数
 *
 * https://leetcode-cn.com/problems/third-maximum-number/description/
 *
 * algorithms
 * Easy (30.80%)
 * Total Accepted:    6.2K
 * Total Submissions: 20K
 * Testcase Example:  '[3,2,1]'
 *
 * 给定一个非空数组，返回此数组中第三大的数。如果不存在，则返回数组中最大的数。要求算法时间复杂度必须是O(n)。
 *
 * 示例 1:
 *
 *
 * 输入: [3, 2, 1]
 *
 * 输出: 1
 *
 * 解释: 第三大的数是 1.
 *
 *
 * 示例 2:
 *
 *
 * 输入: [1, 2]
 *
 * 输出: 2
 *
 * 解释: 第三大的数不存在, 所以返回最大的数 2 .
 *
 *
 * 示例 3:
 *
 *
 * 输入: [2, 2, 3, 1]
 *
 * 输出: 1
 *
 * 解释: 注意，要求返回第三大的数，是指第三大且唯一出现的数。
 * 存在两个值为2的数，它们都排第二。
 *
 *
 */
func ThirdMax(nums []int) int {

	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	var maxM []int

	for n := 0; n < len(nums); n++ {
		max := nums[0]
		index := 0
		for i := range nums {
			if max < nums[i] {
				max = nums[i]
				index = i
			}
		}
		if len(maxM) == 0 {
			maxM = append(maxM, max)
		} else if maxM[len(maxM)-1] != max {
			maxM = append(maxM, max)
		}
		nums[index] = -2147483648

	}

	fmt.Println(maxM)
	if len(maxM) < 3 {
		return maxM[0]
	}
	return maxM[2]
}
