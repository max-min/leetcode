package easy

/*
859. 亲密字符串

给定两个由小写字母构成的字符串 A 和 B ，只要我们可以通过交换 A 中的两个字母得到与 B 相等的结果，就返回 true ；否则返回 false 。



示例 1：

输入： A = "ab", B = "ba"
输出： true
示例 2：

输入： A = "ab", B = "ab"
输出： false
示例 3:

输入： A = "aa", B = "aa"
输出： true
*/

func BuddyStrings(A string, B string) bool {

	if len(A) != len(B) {
		return false
	}

	AT, BT := 0, 0
	diff := 0
	repeat := false
	aMap := make(map[byte]int)
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			diff++
		}
		AT += int(A[i])
		BT += int(B[i])
		if _, exist := aMap[A[i]]; exist {
			repeat = true
		} else {
			aMap[A[i]] = 1
		}
	}
	if AT != BT || (diff == 0 && !repeat) || diff > 2 {
		return false
	}

	return true
}
