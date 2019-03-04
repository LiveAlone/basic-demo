package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	//fmt.Println("first of all content condition")
	//l1 := &ListNode{1, &ListNode{8, nil}}
	//l2 := &ListNode{0, nil}
	//fmt.Println(addTwoNumbers(l1, l2))
	//fmt.Println(lengthOfLongestSubstring("pwwkew"))

	//fmt.Println("0"[0])
	//fmt.Println(myAtoi("-91283472332"))

	//base := []int{2,0}
	//merge(base, 1, []int{1}, 1)
	//fmt.Println(base)

	//fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))

	//fmt.Println(intToRoman(58))

	//fmt.Println(romanToInt("LVIII"))
	//fmt.Println(threeSum([]int{13,14,1,2,-11,-11,-1,5,-1,-11,-9,-12,5,-3,-7,-4,-12,-9,8,-13,-8,2,-6,8,11,7,7,-6,8,-9,0,6,13,-14,-15,9,12,-9,-9,-4,-4,-3,-9,-14,9,-8,-11,13,-10,13,-15,-11,0,-14,5,-4,0,-3,-3,-7,-4,12,14,-14,5,7,10,-5,13,-14,-2,-6,-9,5,-12,7,4,-8,5,1,-10,-3,5,6,-9,-5,9,6,0,14,-15,11,11,6,4,-6,-10,-1,4,-11,-8,-13,-10,-2,-1,-7,-9,10,-7,3,-4,-2,8,-13}))

	//fmt.Println(letterCombinations("23"))
	//head := removeNthFromEnd(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2)
	//for head != nil {
	//	fmt.Println(head.Val)
	//	head = head.Next
	//}

	fmt.Println(generateParenthesis(4))
}

var cache_string_board [][]string

func generateParenthesis(n int) []string {
	cache_string_board = append(cache_string_board, []string{""})
	for len(cache_string_board) <= n {
		currentLength := len(cache_string_board)
		var currResult []string
		for i:=0; i< currentLength; i++ {
			for _, s1 := range cache_string_board[i] {
				for _, s2 := range cache_string_board[currentLength- i - 1] {
					currResult = append(currResult, "(" + s2 + ")" + s1)
				}
			}
		}
		cache_string_board = append(cache_string_board, currResult)
	}
	return cache_string_board[n]
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	size := 0
	curr := head
	for curr!=nil {
		curr = curr.Next
		size++
	}
	curr = head
	dis := size - n
	if dis == 0 {
		return head.Next
	}
	for i:=1; i<dis ; i++ {
		curr = curr.Next
	}
	curr.Next = curr.Next.Next
	return head
}

// phone number calculate
var phone_data = map[uint8]map[int]string{
	'2':{
		1: "a",
		2: "b",
		3: "c",
	},
	'3':{
		1: "d",
		2: "e",
		3: "f",
	},
	'4':{
		1: "g",
		2: "h",
		3: "i",
	},
	'5':{
		1: "j",
		2: "k",
		3: "l",
	},
	'6':{
		1: "m",
		2: "n",
		3: "o",
	},
	'7':{
		1: "p",
		2: "q",
		3: "r",
		4: "s",
	},
	'8':{
		1: "t",
		2: "u",
		3: "v",
	},
	'9':{
		1: "w",
		2: "x",
		3: "y",
		4: "z",
	},
}

func letterCombinations(digits string) []string {
	digits_length := len(digits)
	if digits_length==0 {
		return []string{}
	}

	var result []string

	digits_point := 1
	pos_stack := make([]int, digits_length)
	pos_stack[digits_point-1] = 1

	for digits_point > 0 {
		if pos_stack[digits_point-1] > len(phone_data[digits[digits_point-1]]) {
			digits_point--
			if digits_point == 0 {
				break
			}
			pos_stack[digits_point-1]++
			continue
		}

		if digits_point < digits_length {
			digits_point++
			pos_stack[digits_point-1] = 1
		}else {
			var curr string
			for i:=0; i<digits_length; i++ {
				curr += phone_data[digits[i]][pos_stack[i]]
			}
			result = append(result, curr)
			pos_stack[digits_point-1]++
		}
	}
	return result
}

func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	sort.Ints(nums)
	curr := 0
	dis := int(math.MaxInt32)
	for i := 0; i < length ; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j:=i+1; j < length ; j++ {
			if j > 0 && nums[j] == nums[j-1] && j!=i+1 {
				continue
			}
			for k:=j+1; k < length ; k++ {
				if k > 0 && nums[k] == nums[k-1] && k!=j+1{
					continue
				}
				disCurrent := nums[i] + nums[j] + nums[k] - target
				if disCurrent < 0 {
					disCurrent = -disCurrent
				}
				if disCurrent < dis {
					dis = disCurrent
					curr = nums[i] + nums[j] + nums[k]
				}
			}
		}
	}
	return curr
}

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	length := len(nums)
	sort.Ints(nums)
	for i := 0; i < length ; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j:=i+1; j < length ; j++ {
			if j > 0 && nums[j] == nums[j-1] && j!=i+1 {
				continue
			}
			for k:=j+1; k < length ; k++ {
				if k > 0 && nums[k] == nums[k-1] && k!=j+1{
					continue
				}

				if nums[i] + nums[j] + nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return result
}

func romanToInt(s string) int {
	length := len(s) - 1
	result := 0
	for length >= 0 {
		switch s[length] {
		case 'I':
			result += 1
		case 'V':
			if length > 0 && s[length-1] == 'I' {
				result += 4
				length--
			}else {
				result += 5
			}
		case 'X':
			if length > 0 && s[length-1] == 'I' {
				result += 9
				length--
			}else {
				result += 10
			}
		case 'L':
			if length > 0 && s[length-1] == 'X' {
				result += 40
				length--
			}else {
				result += 50
			}
		case 'C':
			if length > 0 && s[length-1] == 'X' {
				result += 90
				length--
			}else {
				result += 100
			}
		case 'D':
			if length > 0 && s[length-1] == 'C' {
				result += 400
				length--
			}else {
				result += 500
			}
		case 'M':
			if length > 0 && s[length-1] == 'C' {
				result += 900
				length--
			}else {
				result += 1000
			}
		}
		length--
	}
	return result
}


// 计算罗马数字方式

func intToRoman(num int) string {
	var result string
	val := num % 10
	appendResult(val, &result, "I", "V", "X")

	val = (num % 100) / 10
	appendResult(val, &result, "X", "L", "C")

	val = (num % 1000) / 100
	appendResult(val, &result, "C", "D", "M")

	val = num / 1000
	for i:=1; i<=val ; i++ {
		result = "M" + result
	}
	return result
}

func appendResult(val int, result *string, one string, five string, ten string) {
	switch val {
	case 4:
		*result = one + five + *result
	case 9:
		*result = one + ten + *result
	case 5:
		*result = five + *result
	default:
		if val > 4 {
			for i:=6; i<=val ; i++ {
				*result = one + *result
			}
			*result = five + *result
		}else {
			for i:=1; i<=val ; i++ {
				*result = one + *result
			}
		}
	}
}

func maxArea(height []int) int {
	if len(height) == 0 || len(height) == 1 {
		return 0
	}
	l := 0
	r := len(height) - 1
	max := 0
	for l < r {
		if height[l] < height[r] {
			current := height[l] * (r - l)
			if current > max {
				max = current
			}
			l++
		}else {
			current := height[r] * (r - l)
			if current > max {
				max = current
			}
			r--
		}
	}
	return max
}

func climbStairs(n int) int {
	a, b := 1, 2
	for i:=1; i<n; i++ {
		a, b = b, a + b
	}
	return a
}

func deleteDuplicates(head *ListNode) *ListNode {
	currentNode := head
	if currentNode == nil {
		return head
	}
	for {
		if currentNode.Next == nil {
			break
		}
		if currentNode.Val == currentNode.Next.Val {
			currentNode.Next = currentNode.Next.Next
		}else {
			currentNode = currentNode.Next
		}
	}
	return head
}

func merge(nums1 []int, m int, nums2 []int, n int)  {
	cacheArray := make([]int, m)
	for i:=0; i<m ; i++ {
		cacheArray[i] = nums1[i]
	}
	mi := 0
	ni := 0
	pos := 0
	for mi < m || ni < n  {
		if mi == m {
			nums1[pos] = nums2[ni]
			pos++
			ni++
		}else if ni == n {
			nums1[pos] = cacheArray[mi]
			pos++
			mi++
		}else if cacheArray[mi] <= nums2[ni] {
			nums1[pos] = cacheArray[mi]
			pos++
			mi++
		}else {
			nums1[pos] = nums2[ni]
			pos++
			ni++
		}
	}
}

func myAtoi(str string) int {
	var result int64 = 0
	isNegative := false
	hasType := false
	numberStart := false
	for pos:=0; pos<len(str) ; pos++ {
		if str[pos] == 32 {
			if numberStart {
				break
			}
			continue
		}
		if str[pos] == 43 && result == 0 && !hasType{
			if numberStart {
				return 0
			}
			hasType = true
			numberStart = true
			continue
		}
		if str[pos] == 45 && result == 0 && !hasType{
			if numberStart {
				return 0
			}
			isNegative = true
			hasType = true
			numberStart = true
			continue
		}
		if str[pos] - 48 < 10 {
			result = result * 10 + int64(str[pos] - 48)
			numberStart = true
			if result > math.MaxInt32 {
				break
			}
			continue
		}
		break
	}
	if isNegative {
		result = -result
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}
	if result > math.MaxInt32{
		 return math.MaxInt32
	}
	return int(result)
}

// 4 middle value test
func convert(s string, numRows int) string {
	length := len(s)
	result := make([]uint8, length)
	loopSize := 2 * numRows - 2

	for i:=0 ; i<numRows ; i++ {
		for j:=0 ; j + i < length ; j+=loopSize {
			result = append(result, s[j])
			if i != 0 && i != numRows - 1 && j + loopSize - i < length {
				result = append(result, s[j + loopSize - i])
			}
		}
	}
	return string(result)
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
