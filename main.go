package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {

	//509
	//fmt.Println(easy.Fib(8))

	// 500
	//A := []string{"Hello", "Alaska", "Dad", "Peace"}
	//fmt.Println(easy.FindWords(A))

	//811
	//A := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	//fmt.Println(easy.SubdomainVisits(A))

	//682
	//A := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	//fmt.Println(easy.CalPoints(A))

	// 394
	//nums1 := []int{4, 9, 5}
	//nums2 := []int{9, 4, 9, 8, 4}
	//fmt.Println(easy.Intersection(nums1, nums2))

	// 118
	//fmt.Println(easy.Generate(5))

	// 806
	//	A := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	//	S := "bbbcccdddaaa"
	//	fmt.Println(easy.NumberOfLines(A, S))

	//A := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//fmt.Println(easy.RemoveDuplicates(A), A)

	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	fmt.Println(easy.RemoveElement(nums, val), nums)
}
