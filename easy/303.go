package easy

/*
 * @lc app=leetcode.cn id=303 lang=golang
 *
 * [303] 区域和检索 - 数组不可变
 *
 * https://leetcode-cn.com/problems/range-sum-query-immutable/description/
 *
 * algorithms
 * Easy (51.73%)
 * Total Accepted:    8.3K
 * Total Submissions: 16K
 * Testcase Example:  '["NumArray","sumRange","sumRange","sumRange"]\n[[[-2,0,3,-5,2,-1]],[0,2],[2,5],[0,5]]'
 *
 * 给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。
 *
 * 示例：
 *
 * 给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()
 *
 * sumRange(0, 2) -> 1
 * sumRange(2, 5) -> -1
 * sumRange(0, 5) -> -3
 *
 * 说明:
 *
 *
 * 你可以假设数组不可变。
 * 会多次调用 sumRange 方法。
 *
 *
 */
type NumArray struct {
	Nums []int
}

func ConstructorI(nums []int) NumArray {

	n := NumArray{}

	n.Nums = append(n.Nums, nums...)
	return n
}

func (this *NumArray) SumRange(i int, j int) int {

	if i >= len(this.Nums) {
		return 0
	}
	var ret int
	for index := i; index < j+1 && index < len(this.Nums); index++ {
		ret += this.Nums[index]
	}

	return ret
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
