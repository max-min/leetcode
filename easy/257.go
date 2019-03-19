package easy

import "strconv"

/*
 * @lc app=leetcode.cn id=257 lang=golang
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (57.22%)
 * Total Accepted:    6.7K
 * Total Submissions: 11.7K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 *
 * 说明: 叶子节点是指没有子节点的节点。
 *
 * 示例:
 *
 * 输入:
 *
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 *
 * 输出: ["1->2->5", "1->3"]
 *
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
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

func nextTree(ret *[]string, node *TreeNode, st string) {
	st += strconv.Itoa(node.Val) + "->"
	*ret = append(*ret, st)
	if node.Left != nil {
		nextTree(ret, node.Left, st)
	}
	if node.Right != nil {
		nextTree(ret, node.Right, st)
	}

}

func BinaryTreePaths(root *TreeNode) []string {

	if root == nil {
		return []string{""}
	}
	var ret []string

	st := strconv.Itoa(root.Val) + "->"

	ret = append(ret, st)
	if root.Left != nil {
		nextTree(&ret, root.Left, st)
	}

	if root.Right != nil {
		nextTree(&ret, root.Right, st)
	}
	return ret
}
