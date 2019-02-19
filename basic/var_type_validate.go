package basic

import "fmt"

// type 数据类型, error 错误异常断言方式
func main() {

	// i local_string 数据类型支持
	//var i interface{} = "hello world"
	//s := i.(local_string)
	//fmt.Println(s)
	//
	//s, ok := i.(local_string)
	//fmt.Println(s, ok)
	//
	//f, ok := i.(float64)
	//fmt.Println(f, ok)	// 转换失败返回 false, 默认值 float 0 的数值类型
	//
	////f = i.(float64)		// error float64 not support
	////fmt.Println(f)

	// type 数据类型支持
	//do(21)
	//do("hello")
	//do(true)

	// person 数据类型
	//y := Person{"yao", 18}
	//q := Person{"qi", 28}
	//fmt.Println(y, q)
}

type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v user age is %v \n", p.name, p.age)
}

// switch 数据类型判断方式
func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int value type is %v \n", v)
	case string:
		fmt.Printf("local_string value is %v \n", v)
	default:
		fmt.Printf("not known type content is %T \n", v)

	}
}
