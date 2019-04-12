package easy

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=645 lang=golang
 *
 * [645] 错误的集合
 *
 * https://leetcode-cn.com/problems/set-mismatch/description/
 *
 * algorithms
 * Easy (34.86%)
 * Total Accepted:    3.2K
 * Total Submissions: 9.1K
 * Testcase Example:  '[1,2,2,4]'
 *
 * 集合 S 包含从1到 n
 * 的整数。不幸的是，因为数据错误，导致集合里面某一个元素复制了成了集合里面的另外一个元素的值，导致集合丢失了一个整数并且有一个元素重复。
 *
 * 给定一个数组 nums 代表了集合 S 发生错误后的结果。你的任务是首先寻找到重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,2,2,4]
 * 输出: [2,3]
 *
 *
 * 注意:
 *
 *
 * 给定数组的长度范围是 [2, 10000]。
 * 给定的数组是无序的。
 *
 *
 */
func FindErrorNums(nums []int) []int {

	var ret []int
	if len(nums) < 2 {
		return ret
	} else if len(nums) == 2 {
		if nums[0] == nums[1] {
			if nums[0] == 1 {
				ret = append(ret, 1, 2)
			} else {
				ret = append(ret, 2, 1)
			}
		}
		return ret
	}
	sort.Ints(nums)
	fmt.Println(nums)

	exist := false
	first := false
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			ret = append(ret, nums[i])
			first = true
		} else if nums[i]+1 != nums[i+1] {
			exist = true
			first = false
			ret = append(ret, nums[i]+1)
		}
	}
	if first && exist {
		ret[0], ret[1] = ret[1], ret[0]
	}
	if !exist {
		if nums[0] == 1 {
			ret = append(ret, nums[len(nums)-1]+1)
		} else {
			ret = append(ret, 1)
		}

	}
	return ret
}
