package main

import (
	"fmt"
	"github.com/golangBasicDemo/local_string"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(local_string.Revert("hello"))	// lib 对应的引用方式
	fmt.Println(local_string.Revert2("hello"))
}
