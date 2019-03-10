package easy

import "fmt"

func Intersection(nums1 []int, nums2 []int) []int {

	ret := make(map[int]int)
	for i := range nums1 {
		for j := range nums2 {
			if nums1[i] == nums2[j] {
				ret[nums1[i]] = i
				break
			}
		}
	}

	fmt.Println(ret)
	var arr []int
	for k := range ret {
		arr = append(arr, k)
	}

	return arr
}
