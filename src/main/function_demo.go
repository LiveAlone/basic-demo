package main

func main() {

	//fmt.Println("x + y is", add(123, 456))

	//fmt.Println("x + y is", add2(123, 456))

	//a, b := swapString("hello", "world")
	//fmt.Println(a, "123", b)

	//fmt.Println(split(17))
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