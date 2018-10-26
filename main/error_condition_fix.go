package main

import (
	"fmt"
	"math"
	"time"
)

// error 数据类型处理方式
func main() {
	// 基础数据类型
	//if e := run(); e != nil{
	//	fmt.Println(e)
	//}

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

func Sqrt(x float64) (float64, error) {
	if x >=0 {
		return math.Sqrt(x), nil
	}
	return 0, ErrNegativeSqrt(x)
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}


// 自定义异常信息
type MyError struct {
	When time.Time
	What string
}

// 异常函数定义
func (e *MyError) Error() string{
	return fmt.Sprintf("at %v, %s error content \n", e.When, e.What)
}

func run() error{
	return &MyError{time.Now(), "it did not work condition"}
}




