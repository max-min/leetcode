package medium

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//自定义一个类型
type ints []int

func (s ints) Len() int           { return len(s) }
func (s ints) Less(i, j int) bool { return s[i] < s[j] }
func (s ints) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
