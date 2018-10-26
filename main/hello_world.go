package main

import (
	"fmt"
	"math/cmplx"
)

// 基础类型 Demo 支持方式
func main(){

	// 基于 package 导入方式
	//fmt.Println("random int const value is : ", rand.Intn(10))

	// Pi const
	//fmt.Println(math.Pi)

	// function definition
	//fmt.Println(add(1, 100))
	//fmt.Println(addAgain(1, 100))

	// swap content
	//a, b := swap("hello", "world")
	//fmt.Println("content is ", a, b)

	//fmt.Println(split(17))

	// 变量可以声明类型最后
	//var c, python, java bool
	//var i int
	//fmt.Println(i, c, python, java)

	// 变量初始化
	//var i, j = 1, 2
	//var c, python, java = true, false, "no"
	//fmt.Println(i, j, c, python, java)

	// := 可以替换 使用 var 变量方式
	//var i, j = 1, 2
	//k := 10
	//c, python, java := false, true, "no"
	//fmt.Println(i, j, k, c, python, java)

	//fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	//fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	//fmt.Printf("Type: %T Value: %v\n", z, z)

	// 不同的数据类型， 有数据支持
	//var i int
	//var f float64
	//var b bool
	//var s string
	//fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// 通过 T(v) 方式数据转换
	//var x, y = 3, 4
	//var f = math.Sqrt(float64(x * x + y * y))
	//var z = uint(f)
	//fmt.Println(x, y, z)

	// 常量的定义方式
	//const True = 100
	//fmt.Println(True)

	//fmt.Println(nextInt(SMALL))
	////fmt.Println(nextInt(BIG))	// 精度 溢出, 产生对应的数据异常
	//fmt.Println(nextFloat(SMALL))
	//fmt.Println(nextFloat(BIG))
}

func nextInt(x int) int{
	fmt.Println("int get x content is ", x)
	return x * 100 + 1
}

func nextFloat(x float64) float64{
	fmt.Println("float get x content is ", x)
	return x * 100 + 1
}

// 数值常量是一个高精度数值
const (
	BIG = 1 << 100	// int 最大64 位 不会有 100 位
	SMALL = BIG >> 99
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// 返回函数中定义变量， 不推荐 影响可读性
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// 定义多数值类型返回
func swap(x, y string) (string, string){
	return y, x
}

// 函数结构定义
func add(x int, y int) int{
	return x + y
}

// 相同类型可以只申明最后换一个类型
func addAgain(x , y int) int{
	return x + y
}