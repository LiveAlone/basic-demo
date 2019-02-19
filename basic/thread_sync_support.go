package basic

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// 通过 go 方式 go 程并行执行
	//go say("hello")
	//say("world")

	// 通过 信道方式 获取处理的计算结果
	//s := []int{7, 2, 8, -9, 4, 0}
	//c := make(chan int)
	//go sum(s[:len(s)/2], c)
	//go sum(s[len(s)/2:], c)
	//x, y := <-c, <-c // 从 c 中接收
	//fmt.Println(x, y, x + y)

	// 信道支持设置缓冲长度
	//ch := make(chan int, 2)	// 长度小 时候 多线程会阻塞方式
	//ch <- 1
	//ch <- 2
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	//// 信道可以被关闭 v, ok := <-ch false 时候 信道关闭, range 循环等待，等到信道被关闭以后
	//c := make(chan int, 30)
	//go fibonacci(cap(c), c)
	//for i := range c {
	//	fmt.Println(i)
	//}

	// select 方式, 支持 多信道 数据接收方式
	//c := make(chan int)
	//quit := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	//fibonacciSelect(c, quit)

	// 定时方式触发 select 没有注备好执行default
	//tick := time.Tick(1 * time.Second)
	//boom := time.After(10 * time.Second)
	//for {
	//	select {
	//	case <-tick:
	//		fmt.Println("tick.")
	//	case <-boom:
	//		fmt.Println("BOOM!")
	//		return
	//	default:
	//		fmt.Println("    .")
	//		time.Sleep(50 * time.Millisecond)
	//	}
	//}

	// sync Mutex 方式控制并发执行方式
	//c := SafeCounter{v: make(map[local_string]int)}
	//for i := 0; i < 1000; i++ {
	//	go c.Inc("somekey")
	//}
	//
	//time.Sleep(time.Second)
	//fmt.Println(c.Value("somekey"))
}

type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func fibonacciSelect(c, quit chan int){
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		time.Sleep(time.Second)
	}
	close(c)
}

func sum(a []int, c chan int){
	sum := 0
	for i, v := range a {
		fmt.Printf("current sum add is index %v, %v \n", i, v)
		time.Sleep(2 * time.Second)
		sum += v
	}
	c <- sum
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
