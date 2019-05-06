package easy

import (
	"fmt"
	"math"
)

/*
给定一个二叉搜索树的根结点 root, 返回树中任意两节点的差的最小值。

示例：

输入: root = [4,2,6,1,3,null,null]
输出: 1
解释:
注意，root是树结点对象(TreeNode object)，而不是数组。

给定的树 [4,2,6,1,3,null,null] 可表示为下图:

          4
        /   \
      2      6
     / \
    1   3

最小的差值是 1, 它是节点1和节点2的差值, 也是节点3和节点2的差值。
*/

func nextNode(min int, root *TreeNode) int {
	if root.Left != nil {
		l := int(math.Abs(float64(root.Val - root.Left.Val)))
		if l < min {
			min = l
		}
		fmt.Println(root.Val, l, min)
		min = nextNode(min, root.Left)
	}

	if root.Right != nil {
		r := int(math.Abs(float64(root.Val - root.Right.Val)))
		if r < min {
			min = r
		}
		fmt.Println(root.Val, r, min)
		min = nextNode(min, root.Right)
	}
	return min
}

func MinDiffInBST(root *TreeNode) int {

	l, r := -1, -1
	if root.Left != nil {
		l = int(math.Abs(float64(root.Val - root.Left.Val)))
		l = nextNode(l, root.Left)
	}
	if root.Right != nil {
		r = int(math.Abs(float64(root.Val - root.Right.Val)))
		r = nextNode(r, root.Right)
	}
	if l == -1 {
		return r
	}
	if r == -1 {
		return l
	}

	if l > r {
		return r
	}
	return l
}
