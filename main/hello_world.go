package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	fmt.Println()
}

func trailingZeroes(n int) int {

	result := 0
	for n!=0 {
		result += n/5
		n = n / 5
	}
	return result
}
