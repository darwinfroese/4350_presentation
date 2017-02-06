package main

import "fmt"

func sum(c chan int) {
	c <- (1 + 2)
}

func main() {
	c := make(chan int)

	go sum(c)
	go sum(c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
