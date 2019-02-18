package main

func deadPrintContent()  {
	var i = 100
	for {
		i += 1
		i -= 1
	}
}

func main() {
	go deadPrintContent()
	go deadPrintContent()
	deadPrintContent()
}
