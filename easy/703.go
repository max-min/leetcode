package easy

import (
	"sort"
)

/*
设计一个找到数据流中第K大元素的类（class）。注意是排序后的第K大元素，不是第K个不同的元素。

你的 KthLargest 类需要一个同时接收整数 k 和整数数组nums 的构造器，它包含数据流中的初始元素。每次调用 KthLargest.add，返回当前数据流中第K大的元素。

示例:

int k = 3;
int[] arr = [4,5,8,2];
KthLargest kthLargest = new KthLargest(3, arr);
kthLargest.add(3);   // returns 4
kthLargest.add(5);   // returns 5
kthLargest.add(10);  // returns 5
kthLargest.add(9);   // returns 8
kthLargest.add(4);   // returns 8
说明:
你可以假设 nums 的长度≥ k-1 且k ≥ 1。
*/

type KthLargest struct {
	num int
	arr []int
}

func NewConstructor(k int, nums []int) KthLargest {

	kl := KthLargest{
		num: k,
		arr: nums,
	}

	sort.Ints(kl.arr)
	return kl
}

func (this *KthLargest) Add(val int) int {

	i := 0
	for i = 0; i < len(this.arr); i++ {
		if this.arr[i] >= val {
			this.arr = append(this.arr, 0)
			copy(this.arr[i+1:], this.arr[i:])
			this.arr[i] = val
			break
		}
	}
	if i == len(this.arr) {
		this.arr = append(this.arr, val)
	}

	return this.arr[len(this.arr)-this.num]
}
