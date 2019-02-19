package tour

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// 通过 package 包方式, 进行函数数据的导入 导出方式定义

	fmt.Println("current packages content is now")

	fmt.Println("rand int value is ", rand.Intn(100))

	fmt.Println("sqrt num is content ", math.Sqrt(123))

	fmt.Println("Pi content is ", math.Pi)
}