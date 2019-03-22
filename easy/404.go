package easy

/*
 * @lc app=leetcode.cn id=404 lang=golang
 *
 * [404] 左叶子之和
 *
 * https://leetcode-cn.com/problems/sum-of-left-leaves/description/
 *
 * algorithms
 * Easy (48.76%)
 * Total Accepted:    5.1K
 * Total Submissions: 10.3K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 计算给定二叉树的所有左叶子之和。
 *
 * 示例：
 *
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 * 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
 *
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

func leftVal(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Right == nil && root.Left == nil {
		return root.Val
	}

	return leftVal(root.Left) + rightVal(root.Right)
}

func rightVal(root *TreeNode) int {

	if root == nil {
		return 0
	}
	if root.Left == nil {
		return rightVal(root.Right)
	}
	return leftVal(root.Left)
}

func SumOfLeftLeaves(root *TreeNode) int {

	var ret int
	if root == nil {
		return 0
	}
	ret += leftVal(root.Left)
	ret += rightVal(root.Left)
	return ret
}
