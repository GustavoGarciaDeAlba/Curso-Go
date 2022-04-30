package main

import "fmt"

func main() {
	//buffered channel (canal con bufer)
	ca := make(chan int, 1)

	ca <- 42
	ca <- 43

	fmt.Println(<-ca)
}
