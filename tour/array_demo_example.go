package tour

import (
	"fmt"
)

func main() {

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//for i, v := range pow {
	//	fmt.Printf("2**%d = %d\n", i, v)
	//}

	//var s []int
	//printSlice(s)
	//
	//// 添加一个空切片
	//s = append(s, 0)
	//printSlice(s)
	//
	//// 这个切片会按需增长
	//s = append(s, 1)
	//printSlice(s)
	//
	//// 可以一次性添加多个元素
	//s = append(s, 2, 3, 5, 6, 7)
	//printSlice(s)

	//// 创建一个井字板（经典游戏）
	//board := [][]local_string{
	//	[]local_string{"_", "_", "_"},
	//	[]local_string{"_", "_", "_"},
	//	[]local_string{"_", "_", "_"},
	//}
	//
	//// 两个玩家轮流打上 X 和 O
	//board[0][0] = "X"
	//board[2][2] = "O"
	//board[1][2] = "X"
	//board[1][0] = "O"
	//board[0][2] = "X"
	//
	//for i := 0; i < len(board); i++ {
	//	fmt.Printf("%s\n", strings.Join(board[i], " "))
	//}

	// 切片相关
	//a := make([]int, 5)
	//printSlice("a", a)
	//
	//b := make([]int, 0, 5)
	//printSlice("b", b)
	//
	//c := b[:2]
	//printSlice("c", c)
	//
	//d := c[2:5]
	//printSlice("d", d)

	//var s []int
	//fmt.Println(s, len(s), cap(s))
	//if s == nil {
	//	fmt.Println("nil!")
	//}

	//q := []int{2, 3, 5, 7, 11, 13}
	//fmt.Println(q)
	//
	//r := []bool{true, false, true, true, false, true}
	//fmt.Println(r)
	//
	//// 通过定义 struct 方式， 定义数据结构体方式
	//s := []struct {
	//	i int
	//	b bool
	//}{
	//	{2, true},
	//	{3, false},
	//	{5, true},
	//	{7, true},
	//	{11, false},
	//	{13, true},
	//}
	//fmt.Println(s)

	// 数组相关操作
	//var a [2]local_string
	//a[0] = "hello"
	//a[1] = "world"
	//fmt.Println(a)

	//primes := [6]int{1, 2, 3, 4}
	//fmt.Println(primes)

	//primes := [6]int32{2, 3, 5, 7, 11, 13}
	//var s = primes[1:4]
	//fmt.Println(s)
	//fmt.Println(&primes)
	//fmt.Println(&s)
	//s[1] = 6666
	//fmt.Println(primes)	// 切片公用一个数组

	//names := [4]local_string{
	//	"John",
	//	"Paul",
	//	"George",
	//	"Ringo",
	//}
	//fmt.Println(names)
	//
	//a := names[0:2]
	//b := names[1:3]
	//fmt.Println(a, b)
	//
	//b[0] = "XXX"
	//fmt.Println(a, b)
	//fmt.Println(names)

	//s := []int{2, 3, 5, 7, 11, 13}
	//printSlice(s)
	//
	//// 截取切片使其长度为 0
	//s = s[:0]
	//printSlice(s)
	//
	//// 拓展其长度
	//s = s[:4]
	//printSlice(s)
	//
	//// 舍弃前两个值
	//s = s[2:]
	//printSlice(s)
	//
	//s = s[1:3]
	//printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

//func printSlice(s local_string, x []int) {
//	fmt.Printf("%s len=%d cap=%d %v\n",
//		s, len(x), cap(x), x)
//}




















