package easy

/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 *
 * https://leetcode-cn.com/problems/pascals-triangle-ii/description/
 *
 * algorithms
 * Easy (54.12%)
 * Total Accepted:    10.9K
 * Total Submissions: 20.1K
 * Testcase Example:  '3'
 *
 * 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。
 *
 *
 *
 * 在杨辉三角中，每个数是它左上方和右上方的数的和。
 *
 * 示例:
 *
 * 输入: 3
 * 输出: [1,3,3,1]
 *
 *
 * 进阶：
 *
 * 你可以优化你的算法到 O(k) 空间复杂度吗？
 *
 */
func GetRow(rowIndex int) []int {
	var ret []int
	ret = append(ret, 1)
	if rowIndex == 0 {

	} else if rowIndex == 1 {
		ret = append(ret, 1)

	} else {

		last := GetRow(rowIndex - 1)
		for i := 1; i < rowIndex; i++ {
			ret = append(ret, last[i-1]+last[i])
		}
		ret = append(ret, 1)
	}

	return ret
}
