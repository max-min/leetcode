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

	//27
	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	//val := 2
	//fmt.Println(easy.RemoveElement(nums, val), nums)

	//28
	//A := "mississippi"
	//a := "issip"
	//fmt.Println(easy.StrStr(A, a))

	// 35
	//A := []int{1, 3, 5, 6}
	//a := 4
	//fmt.Println(easy.SearchInsert(A, a), A)

	//58
	//A := "A "
	//fmt.Println(easy.LengthOfLastWord(A))

	//A := []int{9, 9}
	//fmt.Println(easy.PlusOne(A))

	//67
	//fmt.Println(easy.AddBinary("110", "001"))

	//69
	//fmt.Println(easy.MySqrt(24))

	//70
	//fmt.Println(easy.ClimbStairs(5))

	//88
	//nums1 := []int{0, 0, 3, 0, 0, 0, 0, 0, 0}
	//m := 3
	//nums2 := []int{-1, 1, 1, 1, 2, 3}
	//n := 6
	//easy.Merge(nums1, m, nums2, n)
	//fmt.Println(nums1)

	//107
	/*
		A := &easy.TreeNode{
			Val: 3,
			Left: &easy.TreeNode{
				Val: 9,
			},
			Right: &easy.TreeNode{
				Val: 20,
				Left: &easy.TreeNode{
					Val: 15,
				},
				Right: &easy.TreeNode{
					Val: 7,
				},
			},
		}
		fmt.Println(easy.LevelOrderBottom(A))
	*/

	// 119
	//fmt.Println(easy.GetRow(3))

	//121
	//A := []int{7, 1, 5, 3, 6, 4}
	//A := []int{1, 2, 3, 4, 5}
	//fmt.Println(easy.MaxProfit2(A))

	//125
	//A := "0P"
	//fmt.Println(easy.IsPalindrome(A))

	//136
	A := []int{4, 4, 3, 1, 3}
	fmt.Println(easy.SingleNumber(A))
}
