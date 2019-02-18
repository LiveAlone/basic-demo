package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	//sum := 0
	//for i := 0 ; i<100 ; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	//sum := 1
	//for ; sum < 5000 ;	  {
	//	sum += sum
	//}
	//fmt.Println(sum)

	//sum := 1
	//for sum < 1000 {
	//	sum += sum
	//}
	//fmt.Print(sum)

	//sum := 1
	//for {
	//	if sum > 5000 {
	//		break
	//	}
	//	sum += sum
	//}
	//fmt.Println(sum)

	//fmt.Println(sqrt(2), sqrt(-4))

	//fmt.Println(
	//	pow(3, 2, 10),
	//	pow(3, 3, 20),
	//)

	//fmt.Println(math.Sqrt(2.0))
	//fmt.Println(sqrt_local(2.0))

	//fmt.Print("Go runs on ")
	//switch_os(runtime.GOOS)

	//switch_func(0)
	//switch_func(1)
	//switch_func(3)
	//switch_func(666)

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func switch_func(i int32) {
	switch i {
	case 0:
		fmt.Println("zero case")
	case one():	// case 0 不会调用
		fmt.Println("one case")
	case 3:
		fmt.Println("three case")
	default:
		fmt.Println("default case")
	}
}

func one() int32{
	fmt.Println("one function called")
	return 1
}

func switch_os(os string) {
	switch os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func sqrt_local(x float64) float64 {
	z := x
	for i := 0; i < 100 ; i++  {
		z -= (z * z - x) / (2 * z)
	}
	return z
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v <= lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

//func sqrt(x float64) string {
//	if x < 0 {
//		return sqrt(-x) + "i"
//	}
//	return fmt.Sprint(math.Sqrt(x))
//}