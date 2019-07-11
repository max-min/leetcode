package medium

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 *
 * https://leetcode-cn.com/problems/3sum/description/
 *
 * algorithms
 * Medium (23.07%)
 * Total Accepted:    68.4K
 * Total Submissions: 296.2K
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0
 * ？找出所有满足条件且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 * 例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 *
 * 满足要求的三元组集合为：
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 *
 *
 */

//自定义一个类型
type ints []int

func (s ints) Len() int           { return len(s) }
func (s ints) Less(i, j int) bool { return s[i] < s[j] }
func (s ints) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func ThreeSum(nums []int) [][]int {
	//var st [][]int
	st := make(map[string][]int)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for m := j + 1; m < len(nums); m++ {
				if nums[i]+nums[j]+nums[m] == 0 {
					arr := []int{nums[i], nums[j], nums[m]}
					sort.Sort(ints(arr))
					k := fmt.Sprintf("%d%d%d", arr[0], arr[1], arr[2])
					st[k] = arr
				}
			}
		}
	}

	var s [][]int
	for _, v := range st {
		s = append(s, v)
	}
	return s
}
