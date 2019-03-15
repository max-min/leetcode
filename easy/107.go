package easy

import "fmt"

/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层次遍历 II
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (59.14%)
 * Total Accepted:    10.9K
 * Total Submissions: 18.5K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
 *
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 * 返回其自底向上的层次遍历为：
 *
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
 * ]
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
func LevelOrderBottom(root *TreeNode) [][]int {

	var ret [][]int
	if root == nil {
		return nil
	}

	ret = append(ret, getArr(ret, root.Left, root.Right))
	top := []int{root.Val}
	ret = append(ret, top)
	return ret
}

func getArr(ret [][]int, left, right *TreeNode) []int {

	if left == nil && right == nil {
		return []int{}
	}
	var inret []int
	var next []int
	if left != nil {
		fmt.Println("left:", left)
		inret = append(inret, left.Val)
		next = append(next, getArr(ret, left.Left, left.Right)...)
	}
	if right != nil {
		fmt.Println("right:", right)
		inret = append(inret, right.Val)
		next = append(next, getArr(ret, right.Left, right.Right)...)
	}

	ret = append(ret, next)
	fmt.Println(ret)
	return inret
}
