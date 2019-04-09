package easy

/*
 * @lc app=leetcode.cn id=551 lang=golang
 *
 * [551] 学生出勤记录 I
 *
 * https://leetcode-cn.com/problems/student-attendance-record-i/description/
 *
 * algorithms
 * Easy (46.92%)
 * Total Accepted:    4.2K
 * Total Submissions: 9K
 * Testcase Example:  '"PPALLP"'
 *
 * 给定一个字符串来代表一个学生的出勤记录，这个记录仅包含以下三个字符：
 *
 *
 * 'A' : Absent，缺勤
 * 'L' : Late，迟到
 * 'P' : Present，到场
 *
 *
 * 如果一个学生的出勤记录中不超过一个'A'(缺勤)并且不超过两个连续的'L'(迟到),那么这个学生会被奖赏。
 *
 * 你需要根据这个学生的出勤记录判断他是否会被奖赏。
 *
 * 示例 1:
 *
 * 输入: "PPALLP"
 * 输出: True
 *
 *
 * 示例 2:
 *
 * 输入: "PPALLL"
 * 输出: False
 *
 *
 */
func CheckRecord(s string) bool {

	aFlag := false
	lTimes := false
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			if aFlag {
				return false
			}
			aFlag = true
		}
		if s[i] == 'L' {
			if lTimes {
				return false
			}
			if i < len(s)-2 {
				if s[i] == s[i+1] && s[i] == s[i+2] {
					lTimes = true
				}
			}
		}
		if aFlag && lTimes {
			return false
		}
	}

	return true
}
