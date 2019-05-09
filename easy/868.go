package easy

import "fmt"

/*
868. 二进制间距
给定一个正整数 N，找到并返回 N 的二进制表示中两个连续的 1 之间的最长距离。

如果没有两个连续的 1，返回 0 。



示例 1：

输入：22
输出：2
解释：
22 的二进制是 0b10110 。
在 22 的二进制表示中，有三个 1，组成两对连续的 1 。
第一对连续的 1 中，两个 1 之间的距离为 2 。
第二对连续的 1 中，两个 1 之间的距离为 1 。
答案取两个距离之中最大的，也就是 2 。
*/

func BinaryGap(N int) int {

	max := 0
	distance := 0
	var binary []byte
	for N > 0 {
		if N%2 == 1 {
			binary = append(binary, '1')
		} else {
			binary = append(binary, '0')
		}

		N /= 2
	}

	b := 0
	fmt.Printf("%s\n", binary)
	for i := range binary {
		if binary[i] == '1' {
			if b > 0 {
				if max < distance {
					max = distance
				}
			}
			distance = 0
			b++
		} else {
			if b > 0 {
				distance++
			}
		}
	}
	if b < 2 {
		return 0
	}
	return max + 1
}
