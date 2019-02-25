package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	//fmt.Println("hello world content")
	//fmt.Println(isPalindrome(10))
	//fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	//fmt.Println(isValid("(]"))
	//fmt.Println(strStr("aa", "aaa"))
	//fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Println(int(math.Sqrt(8)))
}

func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i:=1; i<len(nums); i++ {
		nums[i] = maxInt(nums[i], nums[i] + nums[i-1])
		max = maxInt(max, nums[i])
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func searchInsert(nums []int, target int) int {
	endPos := len(nums)
	pos := 0
	for pos<endPos {
		if nums[pos] >= target {
			return pos
		}
		pos++
	}
	return pos
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i:=0; i < len(haystack) ; i++ {
		if haystack[i] == needle[0] {
			isFinish := true
			for j:=0; j<len(needle); j++ {
				if i + j >= len(haystack) {
					return -1
				}
				if haystack[i+j] != needle[j]{
					isFinish = false
					continue
				}
			}
			if isFinish {
				return i
			}
		}
	}
	return -1
}

func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	p := 0
	i := 0
	for i < len(nums) {
		if nums[i] != val {
			nums[p] = nums[i]
			p+=1
		}
		i++
	}
	return p
}

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	p := 0
	k := 0
	for i, value := range nums {
		if i == 0 {
			k = value
			p += 1
			continue
		}
		if value != k {
			nums[p] = value
			p += 1
			k = value
		}
	}
	return p
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var result *ListNode
	if l1.Val <= l2.Val {
		result = l1
		l1 = l1.Next
	}else {
		result = l2
		l2 = l2.Next
	}

	current := result

	for l1 != nil || l2 != nil {
		if l1 == nil {
			current.Next = l2
			l2 = l2.Next
		}else if l2 == nil {
			current.Next = l1
			l1 = l1.Next
		}else {
			if l1.Val <= l2.Val {
				current.Next = l1
				l1 = l1.Next
			}else {
				current.Next = l2
				l2 = l2.Next
			}
		}
		current = current.Next
	}
	return result
}

func isValid(s string) bool {
	stack := make([]uint8, len(s))
	pos := 0
	for i:=0; i<len(s) ; i++ {
		curr := s[i]
		if curr == 40 || curr == 123 || curr == 91 {
			stack[pos] = curr
			pos += 1
		}else {
			if pos == 0 {
				pos += 1
				break
			}
			toCurr := stack[pos-1]

			switch curr {
			case 93:
				if toCurr != 91 {
					return false
				}
			case 125:
				if toCurr != 123{
					return false
				}
			case 41:
				if toCurr != 40 {
					return false
				}
			}
			pos -= 1
		}
	}
	return pos == 0
}

func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0{
		return ""
	}

	s1 := strs[0]
	commonLength := len(s1)

	for i := 1; i < len(strs) ; i++ {
		currStr := strs[i]
		k := 0
		for k < commonLength && k < len(currStr) {
			if currStr[k] != s1[k] {
				break
			}
			k++
		}
		commonLength = k
		if commonLength == 0 {
			break
		}
	}

	if commonLength == 0 {
		return ""
	}
	return s1[:commonLength]
}

func reverse(x int) int {

	var y int64 = int64(x)

	var revert int64 = 0

	for y != 0{
		revert = revert * 10 + y % 10
		y = y / 10
	}
	if revert <= math.MaxInt32 || revert >= math.MinInt32{
		return int(revert)
	}
	return 0
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := x
	revert := 0
	for y != 0  {
		revert = revert * 10 + y % 10
		y = y / 10
	}
	return revert == x
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		toComplete := target - v
		preIndex, found := m[toComplete]
		if found {
			return []int{preIndex, i}
		}
		m[v] = i
	}
	return []int{}
}
