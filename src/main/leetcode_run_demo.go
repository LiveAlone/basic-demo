package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
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
