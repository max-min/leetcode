package easy

/*
 * @lc app=leetcode.cn id=447 lang=golang
 *
 * [447] 回旋镖的数量
 *
 * https://leetcode-cn.com/problems/number-of-boomerangs/description/
 *
 * algorithms
 * Easy (53.35%)
 * Total Accepted:    3.3K
 * Total Submissions: 6.2K
 * Testcase Example:  '[[0,0],[1,0],[2,0]]'
 *
 * 给定平面上 n 对不同的点，“回旋镖” 是由点表示的元组 (i, j, k) ，其中 i 和 j 之间的距离和 i 和 k
 * 之间的距离相等（需要考虑元组的顺序）。
 *
 * 找到所有回旋镖的数量。你可以假设 n 最大为 500，所有点的坐标在闭区间 [-10000, 10000] 中。
 *
 * 示例:
 *
 *
 * 输入:
 * [[0,0],[1,0],[2,0]]
 *
 * 输出:
 * 2
 *
 * 解释:
 * 两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
 *
 *
 */
func NumberOfBoomerangs(points [][]int) int {

	if len(points) < 3 {
		return 0
	}
	var ret int

	for i := 0; i < len(points); i++ {
		for j := 0; j < len(points); j++ {
			if j == i {
				continue
			}
			for k := 0; k < len(points); k++ {
				if k == j || k == i {
					continue
				}
				x1 := points[i][0] - points[j][0]
				y1 := points[i][1] - points[j][1]
				x2 := points[i][0] - points[k][0]
				y2 := points[i][1] - points[k][1]
				if (x1*x1 + y1*y1) == (x2*x2 + y2*y2) {
					ret++
				}
			}
		}
	}

	return ret
}
