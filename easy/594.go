package easy

import "fmt"

/*
 * @lc app=leetcode.cn id=594 lang=golang
 *
 * [594] 最长和谐子序列
 *
 * https://leetcode-cn.com/problems/longest-harmonious-subsequence/description/
 *
 * algorithms
 * Easy (40.08%)
 * Total Accepted:    2.5K
 * Total Submissions: 6.2K
 * Testcase Example:  '[1,3,2,2,5,2,3,7]'
 *
 * 和谐数组是指一个数组里元素的最大值和最小值之间的差别正好是1。
 *
 * 现在，给定一个整数数组，你需要在所有可能的子序列中找到最长的和谐子序列的长度。
 *
 * 示例 1:
 *
 *
 * 输入: [1,3,2,2,5,2,3,7]
 * 输出: 5
 * 原因: 最长的和谐数组是：[3,2,2,2,3].
 *
 *
 * 说明: 输入的数组长度最大不超过20,000.
 *
 */
func FindLHS(nums []int) int {

	numMap := map[int]int{}
	for _, num := range nums {
		if _, exist := numMap[num]; exist {
			numMap[num]++
		} else {
			numMap[num] = 1
		}
	}
	val, max := 0, 0
	for k, v := range numMap {
		if v > max {
			max = v
			val = k
		}
	}

	fmt.Println(max)
	if max == len(nums) {
		return 0
	} else if max > 1 {
		_, preexist := numMap[val-1]
		_, nextexist := numMap[val+1]

		if preexist && nextexist {
			if numMap[val-1] > numMap[val+1] {
				max += numMap[val-1]
			} else {
				max += numMap[val+1]
			}
		} else if preexist {
			max += numMap[val-1]
		} else {
			max += numMap[val+1]
		}

		return max

	} else {
		for i := range nums {
			for j := 1; j < len(nums); j++ {
				if nums[j]-nums[i] == 1 {
					return 2
				}
			}
		}
		return 1
	}
}
