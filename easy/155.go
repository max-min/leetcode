package easy

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (47.71%)
 * Total Accepted:    18.6K
 * Total Submissions: 38.9K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
 *
 *
 * push(x) -- 将元素 x 推入栈中。
 * pop() -- 删除栈顶的元素。
 * top() -- 获取栈顶元素。
 * getMin() -- 检索栈中的最小元素。
 *
 *
 * 示例:
 *
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 *
 *
 */
type MinStack struct {
	val  int
	next *MinStack
	top  *MinStack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		val:  0,
		next: nil,
		top:  nil,
	}

}

func (this *MinStack) Push(x int) {

	if this.top == nil {
		this.val = x
		this.next = nil
		this.top = this
		return
	}
	tmp := this
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &MinStack{
		val:  x,
		next: nil,
	}
	this.top = tmp.next
}

func (this *MinStack) Pop() {

	if this.top == nil {
		return
	}
	if this.next == nil {
		this.top = nil
		return
	}
	tmp := this.next
	for tmp.next.next != nil {
		tmp = tmp.next
	}
	tmp.next = nil
	this.top = tmp

}

func (this *MinStack) Top() int {
	tmp := this
	for tmp.next != nil {
		tmp = tmp.next
	}
	return tmp.val
}

func (this *MinStack) GetMin() int {

	ret := this.val
	tmp := this.next
	for tmp != nil {
		if tmp.val < ret {
			ret = tmp.val
		}
		tmp = tmp.next
	}
	return ret
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
