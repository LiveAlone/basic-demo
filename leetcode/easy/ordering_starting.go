package main

import "fmt"

func main() {
	//fmt.Println("first of all content condition")
	//l1 := &ListNode{1, &ListNode{8, nil}}
	//l2 := &ListNode{0, nil}
	//fmt.Println(addTwoNumbers(l1, l2))

	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

// 3 longest substring content
func lengthOfLongestSubstring(s string) int {
	strLen := len(s)
	if strLen == 0 {
		return 0
	}
	tagBoard := make([]int, strLen)
	max := 0
	for pos:=0; pos<strLen; pos++ {
		if pos == 0 {
			tagBoard[pos] = 1
			max = 1
			continue
		}
		j:=1
		for j<=tagBoard[pos-1] {
			if s[pos] == s[pos-j] {
				break
			}
			j++
		}
		tagBoard[pos] = j
		if j > max {
			max = j
		}
	}
	return max
}

// 2 add two code
type ListNode struct {
	Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	result := l1
	val := 0
	jinwei := 0
	for {
		val = l1.Val + l2.Val + jinwei
		l1.Val = val % 10
		jinwei = val / 10
		if l1.Next == nil {
			l1.Next = l2.Next
			for jinwei > 0{
				if l1.Next == nil {
					l1.Next = &ListNode{jinwei, nil}
					break
				}
				l1 = l1.Next
				val = l1.Val + jinwei
				l1.Val = val % 10
				jinwei = val / 10
			}
			return result
		}
		if l2.Next == nil {
			for jinwei > 0{
				if l1.Next == nil {
					l1.Next = &ListNode{jinwei, nil}
					break
				}
				l1 = l1.Next
				val = l1.Val + jinwei
				l1.Val = val % 10
				jinwei = val / 10
			}
			return result
		}
		l1 = l1.Next
		l2 = l2.Next
	}
}
