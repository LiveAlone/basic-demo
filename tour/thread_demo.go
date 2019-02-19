package tour

import (
	"sync"
)

//func deadPrintContent()  {
//	var i = 100
//	for {
//		i += 1
//		i -= 1
//	}
//}

//func sum(s []int, c chan int) {
//	sum := 0
//	for _, v := range s{
//		sum += v
//	}
//	c <- sum
//}

type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {

	//c := SafeCounter{v: make(map[local_string]int)}
	//for i := 0; i < 1000; i++ {
	//	go c.Inc("somekey")
	//}
	//time.Sleep(time.Second)
	//fmt.Println(c)
	//fmt.Println(c.Value("somekey"))

	// 多线程 多核 CPU 占用方式
	//go deadPrintContent()
	//go deadPrintContent()
	//deadPrintContent()

	//s := []int{7, 2, 8, -9, 4, 0}
	//c := make(chan int)
	//go sum(s[len(s)/2:], c)	// 启动其他的线程执行
	//go sum(s[:len(s)/2], c)
	//x, y := <-c, <-c
	//fmt.Println(x, y)

	//c := make(chan int, 2)
	//c <- 1
	//c <- 2
	//fmt.Println(<-c)
	//fmt.Println(<-c)

	//c := make(chan int, 10)
	//go fibonacci(cap(c) * 2, c)
	//for v := range c{
	//	fmt.Println(v)
	//}

	//c := make(chan int, 2)
	//quit := make(chan int)
	//go func() {
	//	for i := 0; i < 10 ; i++ {
	//		fmt.Println(<-c)
	//	}
	//	quit <- 0
	//}()
	//fibonacci(c, quit)

	// tick 定时任务,
	//tick := time.Tick(100 * time.Millisecond)
	//boom := time.After(500 * time.Millisecond)
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
}

//func fibonacci(c, quit chan int){
//	x, y := 0, 1
//	for {
//		select {
//		case c <- x:
//			x, y = y, x + y
//		case <-quit:
//			fmt.Println("quit routine content")
//			return
//		}
//	}
//}

//func fibonacci(n int, c chan int) {
//	x, y := 0, 1
//	for i := 0; i < n ; i++  {
//		c <- x
//		x, y = y, x + y
//	}
//	close(c)
//}

