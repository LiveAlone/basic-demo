package main

import (
	"fmt"
)

type T struct {
	name string
	value int
}

func (t T) changeValue() {
	t.value = 100
	fmt.Println("defer t content is ", t)
}

var (
	first = firstValue()
	second = "second"
)

//func init()  {
//	fmt.Println("first value in init is ", first)
//	fmt.Println("second value in init is ", second)
//}

func main() {
	//fmt.Println("hello world 2 condition")
	//t := printHello()
	//fmt.Println("main function t is ", t)
	//fmt.Println(t)

	//a := [3]int{1, 2, 3}
	//fmt.Println(a)
	//changeCopyArrayContent(a)
	//fmt.Println(a)

	//slice := []byte{1, 2, 3}
	//data := []byte{4, 5, 6}
	//result := Append(slice, data)
	//fmt.Println(slice, cap(slice), len(slice))
	//fmt.Println(result, cap(result), len(result))
	//fmt.Println(data, cap(data), len(data))

	//fmt.Printf("Hello %d\n", 23)
	//fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	//fmt.Println("Hello", 23)
	//fmt.Println(fmt.Sprint("Hello ", 23))

	//fmt.Println("main constant value")

	//t := new([]int)
	//fmt.Println(&t)

	//fmt.Println(TwoParamsHandle.handleTow(AddFunction, 1, 2))

	//f := func() {
	//	fmt.Println("test")
	//}
	//f()
	//fp := &f
	//(*fp)()

	// 函数仅仅支持指针方式, 没有变量类似的覆盖方式
	a := adderFn()
	fmt.Println(a())
	fmt.Println(a())

	b := adderFn()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func adderFn() func() int{
	a := 0
	return func() int {
		a ++
		return a
	}
}

func firstValue() string {
	fmt.Println("first value function")
	return "firstValue"
}

func Append(slice, data []byte) []byte {
	l := len(slice)
	if l + len(data) > cap(slice) {  // reallocate
		// Allocate double what's needed, for future growth.
		newSlice := make([]byte, (l+len(data))*2)
		// The copy function is predeclared and works for any slice type.
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:l+len(data)]
	copy(slice[l:], data)
	return slice
}

func changeCopyArrayContent(param [3]int){
	param[0] = 666
}

func printHello() *T {
	t := &T{"yao", 123}
	defer t.changeValue()
	return t
}
