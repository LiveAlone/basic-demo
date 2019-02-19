package tour

import (
	"fmt"
	"math"
)

type Abser interface {
	abs() float64
}

type MyFloat float64

func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex3 struct {
	X, Y float64
}

func (v *Vertex3) abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func (v Vertex3) scale0(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex3) scale(f float64){
	v.X = v.X * f
	v.Y = v.Y * f
}

//func ScaleFunc(v *Vertex3, f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func Abs(v Vertex3) float64{
//	return v.abs()
//}

func main() {
	//v := Vertex3{3, 4}
	//v.scale(66)
	//fmt.Println(v)
	//fmt.Println(v.abs())
	//fmt.Println(Abs(v))

	//f := MyFloat(-math.Sqrt2)
	//fmt.Println(f.abs())

	//v := Vertex3{3, 4}
	//ScaleFunc(&v, 10)
	//fmt.Println(Abs(v))

	//v := Vertex3{3, 4}
	//v.scale0(2)
	//ScaleFunc(&v, 10)
	//fmt.Println(v)
	//
	//p := &Vertex3{4, 3}
	//p.scale0(3)	// 这里只是副本被调用了
	//ScaleFunc(p, 8)
	//fmt.Println(p)

	//var a Abser
	//f := MyFloat(-math.Sqrt2)
	//v := Vertex3{1, 2}
	//a = f
	//fmt.Println(a.abs())
	//a = &v
	//fmt.Println(a.abs())

	//var i I = &T{"listener"}
	//i.M()

	var i I
	fmt.Println(i == nil)
}

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t *T) M() {
	fmt.Println(t.S)
}