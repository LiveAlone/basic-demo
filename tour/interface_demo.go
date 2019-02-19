package tour

import (
	"fmt"
)

func main() {
	//var i interface{}
	//describe(i)
	//i = 42
	//describe(i)
	//i = "hello"
	//describe(i)

	//var i interface{} = "hello"
	//s := i.(local_string)
	//fmt.Println(s)
	//s, ok := i.(local_string)
	//fmt.Println(s, ok)
	//f, ok := i.(float64)
	//fmt.Println(f, ok)
	//f := i.(float64)	// painc
	//fmt.Println(f)

	//do(21)
	//do("hello")
	//do(true)

	//var p1 = Person{"yaoqijun", 18}
	//var p2 = Person{"qijun", 66}
	//fmt.Println(p1, p2)
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func do(i interface{}){
	switch v := i.(type) {
	case int:
		fmt.Println(v * 2)
	case string:
		fmt.Println("local_string value", v)
	default:
		fmt.Printf("not known %T", v)
	}
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

