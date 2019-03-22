package easy

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=401 lang=golang
 *
 * [401] 二进制手表
 *
 * https://leetcode-cn.com/problems/binary-watch/description/
 *
 * algorithms
 * Easy (45.20%)
 * Total Accepted:    2.7K
 * Total Submissions: 5.9K
 * Testcase Example:  '0'
 *
 * 二进制手表顶部有 4 个 LED 代表小时（0-11），底部的 6 个 LED 代表分钟（0-59）。
 *
 * 每个 LED 代表一个 0 或 1，最低位在右侧。
 *
 *
 *
 * 例如，上面的二进制手表读取 “3:25”。
 *
 * 给定一个非负整数 n 代表当前 LED 亮着的数量，返回所有可能的时间。
 *
 * 案例:
 *
 *
 * 输入: n = 1
 * 返回: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16",
 * "0:32"]
 *
 *
 *
 * 注意事项:
 *
 *
 * 输出的顺序没有要求。
 * 小时不会以零开头，比如 “01:00” 是不允许的，应为 “1:00”。
 * 分钟必须由两位数组成，可能会以零开头，比如 “10:2” 是无效的，应为 “10:02”。
 *
 *
 */

func mins(num int) []string {

	var ret []string
	vals := []int{1, 2, 4, 8, 16, 32}
	switch num {
	case 1:

		for i := 0; i < 6; i++ {
			ret = append(ret, fmt.Sprintf("0:%02d", vals[i]))
		}
		//["0:03","0:05","0:06","0:09","0:10","0:12","0:17","0:18","0:20","0:24","0:33","0:34","0:36","0:40","0:48"]
	case 2:
		for i := 0; i < 6; i++ {
			for j := i + 1; j < 6; j++ {
				ret = append(ret, fmt.Sprintf("0:%02d", vals[i]+vals[j]))
			}
		}
	case 3:
		for i := 0; i < 6; i++ {
			for j := i + 1; j < 6; j++ {
				for k := j + 1; k < 6; k++ {
					ret = append(ret, fmt.Sprintf("0:%02d", vals[i]+vals[j]+vals[k]))
				}
			}
		}
	case 4:
		for i := 0; i < 6; i++ {
			for j := i + 1; j < 6; j++ {
				for k := j + 1; k < 6; k++ {
					for m := k + 1; m < 6; m++ {
						if vals[i]+vals[j]+vals[k]+vals[m] < 60 {
							ret = append(ret, fmt.Sprintf("0:%02d", vals[i]+vals[j]+vals[k]+vals[m]))
						}

					}

				}
			}
		}

	case 5:
		for i := 0; i < 6; i++ {
			for j := i + 1; j < 6; j++ {
				for k := j + 1; k < 6; k++ {
					for m := k + 1; m < 6; m++ {
						for n := m + 1; n < 6; n++ {
							if vals[i]+vals[j]+vals[k]+vals[m]+vals[n] < 60 {
								ret = append(ret, fmt.Sprintf("0:%02d", vals[i]+vals[j]+vals[k]+vals[m]+vals[n]))
							}
						}
					}

				}
			}
		}
	}

	return ret
}

func ReadBinaryWatch(num int) []string {

	var ret []string

	if num == 0 {
		ret = append(ret, "0:00")
		return ret
	}
	// hourall
	maxHour := 4
	maxMin := 6

	hourMap := map[int][]string{
		1: {"1:00", "2:00", "4:00", "8:00"},
		2: {"3:00", "5:00", "6:00", "9:00", "10:00"},
		3: {"7:00", "11:00"},
	}
	//fmt.Println(mins(2))
	minuMap := map[int][]string{
		1: mins(1),
		2: mins(2),
		3: mins(3),
		4: mins(4),
		5: mins(5),
	}

	for hourCount := 0; hourCount < maxHour && hourCount <= num; hourCount++ {

		minuCount := num - hourCount
		if minuCount >= maxMin {
			continue
		}
		if hourCount == 0 {
			ret = append(ret, minuMap[minuCount]...)
			continue
		}
		if minuCount == 0 {
			ret = append(ret, hourMap[hourCount]...)
			continue
		}
		for _, h := range hourMap[hourCount] {
			hs := strings.Split(h, ":")
			for _, m := range minuMap[minuCount] {
				ms := strings.Split(m, ":")
				ret = append(ret, hs[0]+":"+ms[1])
			}
		}

	}

	return ret
}
