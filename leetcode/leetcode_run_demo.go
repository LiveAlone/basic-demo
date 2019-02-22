package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	//fmt.Println("hello world content")
	//fmt.Println(isPalindrome(10))
	//fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(isValid("(]"))
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
