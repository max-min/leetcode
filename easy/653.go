package easy

/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入 BST
 *
 * https://leetcode-cn.com/problems/two-sum-iv-input-is-a-bst/description/
 *
 * algorithms
 * Easy (49.84%)
 * Total Accepted:    3.8K
 * Total Submissions: 7.6K
 * Testcase Example:  '[5,3,6,2,4,null,7]\n9'
 *
 * 给定一个二叉搜索树和一个目标结果，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
 *
 * 案例 1:
 *
 *
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 *
 * Target = 9
 *
 * 输出: True
 *
 *
 *
 *
 * 案例 2:
 *
 *
 * 输入:
 * ⁠   5
 * ⁠  / \
 * ⁠ 3   6
 * ⁠/ \   \
 * 2   4   7
 *
 * Target = 28
 *
 * 输出: False
 *
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

func digui1(l, r *TreeNode, k int) bool {
	if l.Val+r.Val == k {
		return true
	}

	left := false
	if l.Left != nil && l.Right != nil {
		left = digui1(l.Left, l.Right, k)
	}
	right := false
	if r.Left != nil && r.Right != nil {
		right = digui1(r.Left, r.Right, k)
	}

	return left || right
}
func FindTarget(root *TreeNode, k int) bool {

	if root == nil {
		return false
	}
	if root.Left != nil && root.Right != nil {
		return digui1(root.Left, root.Right, k)
	}
	return false
}
