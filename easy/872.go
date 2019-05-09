package easy

/*
		        3
	     5             1
	6        2      9     8
		   7   4

请考虑一颗二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个 叶值序列 。

举个例子，如上图所示，给定一颗叶值序列为 (6, 7, 4, 9, 8) 的树。

如果有两颗二叉树的叶值序列是相同，那么我们就认为它们是 叶相似 的。

如果给定的两个头结点分别为 root1 和 root2 的树是叶相似的，则返回 true；否则返回 false 。



提示：

给定的两颗树可能会有 1 到 100 个结点。
*/

func noteLeaves(root *TreeNode) []int {
	var ret []int
	if root.Left == nil && root.Right == nil {
		return append(ret, root.Val)
	}
	if root.Left != nil {
		ret = append(ret, noteLeaves(root.Left)...)
	}
	if root.Right != nil {
		ret = append(ret, noteLeaves(root.Right)...)
	}
	return ret
}

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {

	s1 := noteLeaves(root1)
	s2 := noteLeaves(root2)

	//fmt.Println(s1, s2)
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
