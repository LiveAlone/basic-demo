package tour

import "fmt"

func main() {

	//fmt.Println("x + y is", add(123, 456))

	//fmt.Println("x + y is", add2(123, 456))

	//a, b := swapString("hello", "world")
	//fmt.Println(a, "123", b)

	//fmt.Println(split(17))

	//hypot := func(x, y float64) float64 {
	//	return math.Sqrt(x*x + y*y)
	//}
	//fmt.Println(hypot(5, 12))
	//fmt.Println(compute(hypot))
	//fmt.Println(compute(math.Pow))

	//pos, neg := adder(), adder()
	//for i := 0 ; i < 10 ; i++ {
	//	fmt.Println(pos(2 * i))
	//	fmt.Println(neg(-i))
	//}

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		c := a
		a, b = b, a + b
		return c
	}
}

func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	//x = 666
	//return 123, 456
	return
}

func swapString(x, y string) (string, string) {
	return y, x
}

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int{
	return x + y
}