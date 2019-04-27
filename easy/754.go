package easy

/*
在一根无限长的数轴上，你站在0的位置。终点在target的位置。

每次你可以选择向左或向右移动。第 n 次移动（从 1 开始），可以走 n 步。

返回到达终点需要的最小移动次数。

示例 1:

输入: target = 3
输出: 2
解释:
第一次移动，从 0 到 1 。
第二次移动，从 1 到 3 。
示例 2:

输入: target = 2
输出: 3
解释:
第一次移动，从 0 到 1 。
第二次移动，从 1 到 -1 。
第三次移动，从 -1 到 2 。
*/

/*
func leftgo(stepResult, steps, target int) int {
	stepResult -= steps
	if target == stepResult {
		return steps
	}
	l := leftgo(stepResult, steps+1, target)
	r := rightgo(stepResult, steps+1, target)
	if l < r {
		return l
	}
	return r

}

func rightgo(stepResult, steps, target int) int {
	stepResult += steps
	if target == stepResult {
		return steps
	}
	l := leftgo(stepResult, steps+1, target)
	r := rightgo(stepResult, steps+1, target)
	if l < r {
		return l
	}
	return r

}

func ReachNumber(target int) int {

	step := 0
	l := leftgo(step, 1, target)
	r := rightgo(step, 1, target)
	if l < r {
		return l
	}
	return r
}
*/
func ReachNumber(target int) int {

	if target < 0 {
		target *= -1
	}

	step := 0
	stepTotal := 0
	for stepTotal <= target {
		step++
		stepTotal += step
	}

	if (stepTotal-target)%2 == 0 {
		return step
	}

	if step%2 == 0 {
		return step + 1
	} else {
		return step + 2
	}
}
