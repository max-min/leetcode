package easy

/*
 * @lc app=leetcode.cn id=101 lang=golang
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (45.29%)
 * Total Accepted:    23.9K
 * Total Submissions: 52.6K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给定一个二叉树，检查它是否是镜像对称的。
 *
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 *
 *
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 *
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 *
 *
 * 说明:
 *
 * 如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
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
func IsSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	if root.Left == nil || root.Right == nil {
		return false
	}

	if root.Right.Val == root.Left.Val {
		return twiceNode(root.Left.Left, root.Right.Right) && twiceNode(root.Left.Right, root.Right.Left)
	}
	return false
}

func twiceNode(left *TreeNode, right *TreeNode) bool {

	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	if left.Val == right.Val {
		return twiceNode(left.Left, right.Right) && twiceNode(left.Right, right.Left)
	}
	return false
}
