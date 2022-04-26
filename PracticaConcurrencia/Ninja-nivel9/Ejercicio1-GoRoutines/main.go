package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	var waitG sync.WaitGroup
	waitG.Add(2)

	go func() {
		fmt.Println("a")
		waitG.Done()
	}()

	go func() {
		fmt.Println("todos")
		waitG.Done()
	}()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	fmt.Println("Hola")
	waitG.Wait()
}
