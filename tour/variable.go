package tour

import "fmt"

//var c, python, java bool

//test := 123	// error
//var test = 123
//
//var (
//	ToBe   bool       = false
//	MaxInt uint64     = 1<<64 - 1
//	y float64 = math.Sqrt(123)
//	z      complex128 = cmplx.Sqrt(-5 + 12i)
//)

//const Pi = 3.14

const (
	// 将 1 左移 100 位来创建一个非常大的数字
	// 即这个数的二进制是 1 后面跟着 100 个 0
	Big = 1 << 100
	// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	//const World = "世界"
	//fmt.Println("Hello", World)
	//fmt.Println("Happy", Pi, "Day")
	//
	//const Truth = true
	//fmt.Println("Go rules?", Truth)

	//var x, y int = 7, 11
	//var f float64 = math.Sqrt(float64(x*x + y*y))
	//var z uint = uint(f)
	//fmt.Println(x, y, f, z)

	//var i int
	//fmt.Println(i, c, python, java)

	//var i, j = 1, 2
	//fmt.Println(i, j)

	//var a, b, c = true, 0, "hello world"
	//fmt.Println(a, b, c)

	//var i, j int = 1, 2
	//k := 3
	//c, python, java := true, false, "no!"
	//
	//fmt.Println(i, j, k, c, python, java)

	//fmt.Println(test)

	//fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	//fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	//fmt.Printf("Type: %T Value: %v\n", y, y)
	//fmt.Printf("Type: %T Value: %v\n", z, z)
}
