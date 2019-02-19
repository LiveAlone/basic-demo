package basic

import (
	"fmt"
	"math"
)

// function 基础功能
func main() {

	//v := Vertex2{3, 4}
	//fmt.Println(v.Abs())

	// 上面定义类似函数方法, 那么静态方法呢， 直接声明就好了
	//v := Vertex2{3, 4}
	//fmt.Println(Abs(v))

	// 对于内建类型 不能动态 声明方法，
	//f := MyFloat(-math.Sqrt2)
	//fmt.Println(f.Abs())

	// struct 对于 参数 也是 拷贝的副本， 只有地址才能修改对应的数值 ！！
	//v := Vertex2{3, 4}
	//v.Scale(10)
	//fmt.Println(v.Abs())

	// 使用大对象时候 使用指针， 减少函数调用， 数据的复制！！！！

	// 接口继承使用方式
	//var a Abser
	//f := MyFloat(-math.Sqrt2)
	//v := Vertex2{3, 4}
	//
	//a = f
	////a = v	// 错误 &v 指针实现的接口
	//a = &v
	//fmt.Println(a.Abs())

	// go 类型断言支持， 减少数据的支持
	//var i interface{} = "hello"
	//
	//s := i.(local_string)
	//fmt.Println(s)
	//
	//s, ok := i.(local_string)
	//fmt.Println(s, ok)
	//
	//f, ok := i.(float64)	// 兼容类型转换， 判断结果正确性
	//fmt.Println(f, ok)
	//
	////f = i.(float64) // panic
	////fmt.Println(f)

	//do(21)
	//do("hello")
	//do(true)

	// nil 数据类型的操作方式
	//var i I
	//fmt.Printf("(%v %T)", i, i)
	////i.M()	// error nil 并没有异常信息

	// 空接口 数据类型
	var i interface{}
	describeInterface(i)

	i = 42
	describeInterface(i)

	i = "hello"
	describeInterface(i)
}

// 函数指定接口数据类型
func describeInterface(i interface{}){
	fmt.Printf("(%v, %T) \n", i, i)
}

type I interface {
	M()
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// 定义接口
type Abser interface {
	Abs() float64
}

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64{
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func Abs(v Vertex2) float64{
	return math.Sqrt(v.Y * v.Y + v.X * v.X)
}

func (v *Vertex2) Abs() float64{
	return math.Sqrt(v.Y * v.Y + v.X * v.X)
}

type Vertex2 struct {
	X, Y float64
}