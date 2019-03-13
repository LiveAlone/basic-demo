package basic

import (
	"fmt"
	"math"
	"time"
)

// 运算执行流程方式
func main() {

	// for sum 100
	//sum := 0
	//for i:=0; i<100; i++ {
	//	sum += i
	//}
	//fmt.Println("sum amount content is ", sum)

	// 允许取出 ; 分号
	//sum := 2
	//for ; sum < 1000; {
	//	sum += sum
	//}
	//fmt.Println(sum)
	//
	//sum := 2
	//for sum < 199 {
	//	sum += sum
	//}
	//fmt.Println(sum)

	// 无条件 不需要 for true
	//sum := 2
	//for {
	//	if sum > 100 {
	//		break
	//	}
	//	sum += 2
	//}
	//fmt.Println(sum)

	// if 类似 for 执行前加 语句
	//fmt.Println(pow(3, 2, 10))
	//fmt.Println(pow(3, 3, 10))

	////switch 默认break 除非 fallthrough
	//fmt.Println("Go run content ")
	//switch os:=runtime.GOOS; os{
	//case "darwin":
	//	fmt.Println("OS X.")
	//	fallthrough
	//case "linux":
	//	fmt.Println("Linux.")
	//default:
	//	// freebsd, openbsd,
	//	// plan9, windows...
	//	fmt.Printf("%s.", os)
	//}

	// switch case i 可以指定方法， 判断方法执行
	//fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// switch 支持没有条件的判断方式, t.Hour() 最多两次掉用方式
	//t := time.Now()
	//fmt.Println("current hour is", t.Hour())
	//switch {
	//case t.Hour() < 12:
	//	fmt.Println("Good morning!")
	//case t.Hour() < 17:
	//	fmt.Println("Good afternoon.")
	//default:
	//	fmt.Println("Good evening.")
	//}

	// defer 函数调用完成后， 执行调用返回结果
	//fmt.Println("start before defer")
	//def()
	//fmt.Println("after end defer")

	//fmt.Println("counting")
	//for i := 0; i < 10; i++ {
	//	defer fmt.Println(i)
	//}
	//fmt.Println("done")
}

func def() {

	defer fmt.Println("world")

	fmt.Println("hello")

	//return  // return 后 ！！ 就不会输出

	defer fmt.Println("!!")	// 这个会比 world 先执行 stack 战方式执行
}

func pow(x, y, lim float64) float64{
	if v := math.Pow(x, y); v < lim {
		return v
	}else {
		fmt.Println("exceed content lim", lim)
	}
	// 现在不能使用 v
	//return v
	return lim
}
