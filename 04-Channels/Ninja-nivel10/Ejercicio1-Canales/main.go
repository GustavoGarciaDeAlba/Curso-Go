package main

import (
	"fmt"
)

func main() {

	//Canal sin buffer
	//c := make(chan int)

	//Canal con buffer
	c := make(chan int, 1)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}
