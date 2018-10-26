package main

// 函数参数方式 可执行配置方式
func main() {

	//hypot := func(x, y float64) float64 {
	//	return math.Sqrt(x * x + y * y)
	//}
	//
	//fmt.Println(hypot(5, 12))
	//fmt.Println(compute(hypot))
	//fmt.Println(compute(math.Pow))

	//pos, neg := addOrder(), addOrder()
	//for i:=0; i<10; i++ {
	//	fmt.Println(
	//		pos(i),
	//		neg(2 * -i),
	//		)
	//}


}

// 函数结果返回支持
func addOrder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

// 传递函数方式， 进行函数传递搜索方式
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
