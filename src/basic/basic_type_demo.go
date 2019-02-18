package basic

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome test go condition")

	fmt.Println("current time is : ", time.Now().AddDate(0, 0, 1))

	fmt.Println("sub value is : ", int64(time.Now().AddDate(0, 0, 1).Sub(time.Now())) / 1000000000 / 3600)
}