package tour

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v1 := Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 := Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 := Vertex{}      // X:0 Y:0
	p := &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
	fmt.Println(v1, p, v2, v3)

	//i, j := 42, 2701
	//
	//p := &i
	//fmt.Println(p)
	//fmt.Println(*p)
	//
	//*p = 21
	//fmt.Println(i)
	//
	//p = &j
	//*p = *p / 37
	//fmt.Println(j)

	//fmt.Println(Vertex{1, 2})

	//v := Vertex{1,2}
	//v.Y = 123
	//fmt.Println(v)

	//v := Vertex{123, 456}
	//p := &v
	//fmt.Println(p)
	//fmt.Println((*p).Y)
	//fmt.Println(p.X)
	//p.X = 1e9
	//fmt.Println(v)

	//v := Vertex{123, 456}
	//fmt.Println(&v.X)
	//v.X = 666
	//fmt.Println(&v.X)
	//fmt.Println(v)
}
