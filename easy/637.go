package easy

/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
 *
 * https://leetcode-cn.com/problems/average-of-levels-in-binary-tree/description/
 *
 * algorithms
 * Easy (58.01%)
 * Total Accepted:    4.7K
 * Total Submissions: 8.1K
 * Testcase Example:  '[3,9,20,15,7]'
 *
 * 给定一个非空二叉树, 返回一个由每层节点平均值组成的数组.
 *
 * 示例 1:
 *
 * 输入:
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 输出: [3, 14.5, 11]
 * 解释:
 * 第0层的平均值是 3,  第1层是 14.5, 第2层是 11. 因此返回 [3, 14.5, 11].
 *
 *
 * 注意：
 *
 *
 * 节点值的范围在32位有符号整数范围内。
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func digui(ret *[]float64, l, r *TreeNode) {

	if l == nil && r == nil {
		return
	}
	num := 0
	val := 0
	if l != nil {
		num++
		val += l.Val
	}
	if r != nil {
		num++
		val += r.Val
	}
	if num > 0 {
		*ret = append(*ret, float64(val)/float64(num))
	}
	if l != nil && r != nil {
		digui(ret, l.Left, l.Right)
	}
	if r != nil {
		digui(ret, r.Left, r.Right)
	}

}
func averageOfLevels(root *TreeNode) []float64 {

	var ret []float64
	if root != nil {
		ret = append(ret, float64(root.Val))
	}
	digui(&ret, root.Left, root.Right)
	return ret
}
