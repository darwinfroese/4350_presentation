package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go add()
	}
}

func add() {
	fmt.Printf("1 + 2 = %d\n", 1+2)
}
