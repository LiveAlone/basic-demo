package basic

import (
	"fmt"
)

// 指针 结构体定时修改方式
func main() {

	//i, j := 3, 2701
	//p := &i
	//fmt.Println(*p)
	//*p = 21
	//fmt.Println(p)
	//fmt.Println(i)	// 指向地址空间数据已经被修改了
	//
	//p = &j
	//*p = *p / 37
	//fmt.Println(j)

	// 结构体类型支持
	//fmt.Println(Vertex{1, 2})
	//a, b := Vertex{1, 2}, Vertex{3, 4}
	//fmt.Println(a, b)
	////fmt.Print(&a, &b)	// 好像不支持

	//v := Vertex{1, 2}
	//p := &v
	//fmt.Println(p)	// 估计不支持地址运算方式
	//// 不用 (*p).Y
	//p.Y = 666
	//v.X = 200
	//fmt.Println(v)

	//fmt.Println(v1, p, v2, v3)

	// 数组支持方式
	//var a [2]local_string
	//a[0] = "Hello"
	//a[1] = "World"
	//fmt.Println(a[0], a[1])
	//fmt.Println(a)
	//
	//primes := [6]int{2, 3, 11, 13}
	//fmt.Println(primes)

	// 支持切片方式
	//primes := [6]int{2, 3, 5, 7, 11, 13}
	//var s []int = primes[1:4]
	//fmt.Println(s)

	// 数组切片 底层共享
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

	// 不同的数据类型 支持 数组
	//q:= []int{1,2,3,4,5,6}
	//fmt.Println(q)
	//
	//r := []bool{false, false, true, false}
	//fmt.Println(r)
	//
	//s := []struct{	// struct 定义和 数组的生成同时
	//	i int
	//	b bool
	//}{
	//	{2, false},
	//	{3, true},
	//	{4, false},
	//}
	//fmt.Println(s)

	// 切片不同的视角支持
	//s := []int{1, 3, 5, 7, 9, 11, 13, 15}
	//printSlice(s)
	//printSlice(s[:0])
	//printSlice(s[:4])
	//printSlice(s[2:])

	// nil
	//var s []int
	//fmt.Println(s, len(s), cap(s))
	//if s == nil {
	//	fmt.Println("s content is nil")
	//}

	// 数组初始化构建方式
	//var s1 [4]int
	//fmt.Println(s1)
	//var s2 = make([]int, 4)
	//fmt.Println(s2)
	//b := make([]int, 0, 5)
	//printSlice(b)
	//c := b[:2]
	//printSlice(c)
	//d := c[2:5]
	//printSlice(d)

	// 二维数据切片支持
	//board := [][]local_string{
	//	[]local_string{"_", "_", "_"},
	//	[]local_string{"_", "_", "_"},
	//	[]local_string{"_", "_", "_"},
	//}
	//// The players take turns.
	//board[0][0] = "X"
	//board[2][2] = "O"
	//board[1][2] = "X"
	//board[1][0] = "O"
	//board[0][2] = "X"
	//for i := 0; i < len(board); i++ {
	//	fmt.Printf("%s\n", strings.Join(board[i], " "))
	//}

	// array to resize
	//var s []int
	//printSlice(s)
	//
	//s1 := append(s, 1)
	//printSlice(s)
	//printSlice(s1)
	//
	//s2 := append(s1, 2)
	//printSlice(s)
	//printSlice(s1)
	//printSlice(s2)
	//
	//s3 := append(s2, 3)
	//printSlice(s)
	//printSlice(s1)
	//printSlice(s2)
	//printSlice(s3)

	// range 函数方式 支持数组遍历 index value condition, 当然可以忽略方式
	//var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//for i, v := range pow {
	//	fmt.Printf("2**%d = %d\n", i, v)
	//}
	//pow := make([]int, 10)
	//for i := range pow {
	//	pow[i] = 1 << uint(i) // == 2**i
	//}
	//for _, value := range pow {
	//	fmt.Printf("%d\n", value)
	//}

	// map 数据映射结构
	////var m map[local_string]Vertex // current map is nil， 需要通过 make 初始化内存
	////fmt.Println(m == nil)
	//m := make(map[local_string]Vertex)
	//fmt.Println(m == nil)
	//fmt.Println(m)
	//fmt.Println("len map is ", len(m))
	//m["Bell Labs"] = Vertex{
	//	68433, -39967,
	//}
	//fmt.Println(m)

	// map 初始化方式 可折叠方式， 对于 struct 相同的数据类型支持
	//m := map[local_string]Vertex{
	//	"yao": Vertex{1,2},
	//	"qi": Vertex{2, 3},
	//}
	//m := map[local_string]Vertex{
	//	"yao":{1, 2},
	//	"qi":{2,3},
	//}
	//fmt.Println(m)

	// 获取元素方式， 判断元素存在方式
	//m := make(map[local_string]int)
	//m["Answer"] = 42
	//fmt.Println("The value:", m["Answer"])
	//m["Answer"] = 48
	//fmt.Println("The value:", m["Answer"])
	//delete(m, "Answer")
	//fmt.Println("The value:", m["Answer"])	 // no null point exception 0 default condition
	//v, ok := m["Answer"]
	//fmt.Println("The value:", v, "Present?", ok)
}

func printSlice(a []int) {
	fmt.Printf("len is %d, cap is %d, content is %v \n", len(a), cap(a), a)
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

type Vertex struct {
	X int
	Y int
}
