package easy

import "fmt"

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 *
 * https://leetcode-cn.com/problems/merge-sorted-array/description/
 *
 * algorithms
 * Easy (43.05%)
 * Total Accepted:    31.8K
 * Total Submissions: 73.9K
 * Testcase Example:  '[1,2,3,0,0,0]\n3\n[2,5,6]\n3'
 *
 * 给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。
 *
 * 说明:
 *
 *
 * 初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
 * 你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
 *
 *
 * 示例:
 *
 * 输入:
 * nums1 = [1,2,3,0,0,0], m = 3
 * nums2 = [2,5,6],       n = 3
 *
 * 输出: [1,2,2,3,5,6]
 *
 */
func Merge(nums1 []int, m int, nums2 []int, n int) {

	var ret []int
	i, j := 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			ret = append(ret, nums1[i])
			i++
		} else {
			ret = append(ret, nums2[j])
			j++
		}
	}
	fmt.Println("i=", i, "j=", j, ret)
	if i < m {
		ret = append(ret, nums1[i:m]...)
	}
	if j < n {
		ret = append(ret, nums2[j:]...)
	}
	for i := range ret {
		nums1[i] = ret[i]
	}
}
