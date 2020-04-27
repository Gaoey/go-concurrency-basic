package main

import "fmt"

var (
	x = 0
)

func main() {
	ch := make(chan byte, 1)
	go add("First", ch)
	go add("Second", ch)
	// flow out
	<-ch
}

func add(who string, ch chan byte) {
	for i := 0; i <= 10; i++ {
		x += i
		fmt.Printf("%s | i = %d, x = %d\n", who, i, x)
	}
	// flow in
	ch <- 1
}
