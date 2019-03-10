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

	A := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	fmt.Println(easy.CalPoints(A))
}
