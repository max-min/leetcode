package easy

/*
896. 单调数列

如果数组是单调递增或单调递减的，那么它是单调的。

如果对于所有 i <= j，A[i] <= A[j]，那么数组 A 是单调递增的。 如果对于所有 i <= j，A[i]> = A[j]，那么数组 A 是单调递减的。

当给定的数组 A 是单调数组时返回 true，否则返回 false。
*/

func IsMonotonic(A []int) bool {

	max := false
	min := false
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			if min {
				return false
			}
			max = true
		} else if A[i] < A[i+1] {
			if max {
				return false
			}
			min = true
		}
	}
	if max && min {
		return false
	}
	return true
}
