package easy

import "fmt"

/*
 * @lc app=leetcode.cn id=448 lang=golang
 *
 * [448] 找到所有数组中消失的数字
 *
 * https://leetcode-cn.com/problems/find-all-numbers-disappeared-in-an-array/description/
 *
 * algorithms
 * Easy (46.85%)
 * Total Accepted:    5.8K
 * Total Submissions: 12.4K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * 给定一个范围在  1 ≤ a[i] ≤ n ( n = 数组大小 ) 的 整型数组，数组中的元素一些出现了两次，另一些只出现一次。
 *
 * 找到所有在 [1, n] 范围之间没有出现在数组中的数字。
 *
 * 您能在不使用额外空间且时间复杂度为O(n)的情况下完成这个任务吗? 你可以假定返回的数组不算在额外空间内。
 *
 * 示例:
 *
 *
 * 输入:
 * [4,3,2,7,8,2,3,1]
 *
 * 输出:
 * [5,6]
 *
 *
 */
func FindDisappearedNumbers(nums []int) []int {

	var ret []int
	if len(nums) == 0 {
		return ret
	}
	max := 0
	for _, num := range nums {
		if max < num {
			max = num
		}
	}
	numMap := map[int]int{}
	repeats := 0
	for i := 1; i <= max; i++ {
		exist := false
		for _, num := range nums {
			if num == i {
				exist = true
				if _, exist := numMap[num]; exist {
					numMap[num]++
					repeats++
				} else {
					numMap[num] = 1
				}
			}
		}
		if !exist {
			ret = append(ret, i)
		}
	}
	fmt.Println(ret)
	for _, v := range numMap {
		for v > 1 {
			if len(ret) < repeats {
				max++
				ret = append(ret, max)
				v--
			} else {
				break
			}
		}
	}

	return ret
}
