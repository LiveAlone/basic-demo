package basic

import (
	"fmt"
	"io"
	"strings"
)

// local_string file content reader
func main() {
	r := strings.NewReader("hello world")
	b := make([]byte, 2)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			// 文件读取完成 error = EOF
			break
		}
	}
}
